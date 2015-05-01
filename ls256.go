// ls256.go (c) 2013-2014 David Rook - all rights reserved
//
// list256 dir
// typical output:     6003568 | 9df360b8ec6a58f6cb410303c7794d98526cbfe2b11be7a34754515d0fcb21bb | 2013-07-03:09_03_57 | /home/mdr/Desktop/MYGO/src/vwar3/vwar3
//
// limitation - gathers filenames as an argument list before procesing.  Can exhaust memory if too many files
//
package main

// BUG(mdr): count non-regular files and report
// BUG(mdr): better report of bad date files - or save in logfile?
// BUG(mdr): we're ignoring this for now 	-links  directory link count in output	
//	flag.BoolVar(&g_LinksFlag, "links", false, "Include directory link counts in output")

import (
	// go stdlib 1.4.2 pkgs
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	// local pkgs
	"github.com/hotei/mdr"
)

const (
	G_version = "ls256.go   (c) 2013-2015 David Rook version 0.0.5"
)

var (
	flagCPU       int // cpus requested on command line
	g_argList     []string
	g_verboseFlag bool
	g_badDateCt   int
	g_noSHAFlag   bool
	doExtFilter   bool
	g_LinksFlag   bool
	g_extFilter   string
	g_tmMutex     sync.Mutex
	outPut        sync.WaitGroup
	loop          sync.WaitGroup
	nCPU          int = 8
)

func init() {
	flag.BoolVar(&g_verboseFlag, "verbose", false, "Verbose messages")
	flag.BoolVar(&g_noSHAFlag, "nosha", false, "skip sha computation")
	flag.IntVar(&flagCPU, "cpu", 0, "Number of CPU cores to use(default is all available)")
	flag.StringVar(&g_extFilter, "ext", "", "Extension to match (default is all)")
	g_argList = make([]string, 0, 10)
}

func usage() {
	txt := `
ls256 [OPTIONS] path

OPTIONS:
	-verbose               more intermediate messages
	-nosha                 dont compute SHA256  (still does size date and name)
	-cpu=n                 limit operation to n CPUs (default is all available)
	-ext=".ext"            limit to files with this extension (default is all)
	`
	fmt.Printf("%s\n", txt)
	os.Exit(0)
}

type digestType struct {
	pathname   string
	dig256     string
	fileLength int64
	fileDate   time.Time
}

func lineOut(c chan digestType) {
	var d digestType
	outPut.Add(1) // waitgroup
	defer outPut.Done()
	for {
		d = <-c
		if d.pathname == "" {
			return
		}
		fmt.Printf("%12d | %64s | %s | %s \n",
			d.fileLength, d.dig256, d.fileDate.Format("2006-01-02:15_04_05"), d.pathname)
	}
}

// returns filepath.SkipDir on encountering .gvfs indicating a local mount of a remote file system
// might want a flag option to switch this behavior
//
func CheckPath(pathname string, info os.FileInfo, err error) error {
	if info == nil {
		fmt.Printf("WARNING --->  no stat info available for %s\n", pathname)
		return nil
	}
	if info.IsDir() {
		Verbose.Printf("Checking path %s\n", pathname)
		Verbose.Printf("dir = %v\n", pathname)
		if strings.Contains(pathname, "/.gvfs/") {
			fmt.Printf("found a .gvfs dir (remote filesystem) - skipping it\n")
			return filepath.SkipDir
		}
		if g_LinksFlag {
			nlinks, err := mdr.FileLinkCt(pathname)
			if err != nil {
				fmt.Printf("# err %v getting link ct for %s\n", err, pathname)
				return nil
			}
			fmt.Printf("# dirLinks | %d | %s\n", nlinks, pathname)
		}
		mdr.Spinner()
	} else { // regular file
		fmode := info.Mode()
		if fmode.IsRegular() == false {
			Verbose.Printf("non-regular file skipped -> %s\n", pathname)
			// BUG(mdr): save skipped files in a list for appending?
			return nil
		}
		if doExtFilter {
			ext := strings.ToLower(path.Ext(pathname))
			if ext != g_extFilter {
				return nil // not right extension
			}
		}
		Verbose.Printf("%10d %s\n", info.Size(), pathname)
		g_argList = append(g_argList, pathname)
	}
	return nil
}

func flagSetup() {
	// -ext=".mp3" for example
	if len(g_extFilter) > 0 {
		doExtFilter = true
		g_extFilter = strings.ToLower(g_extFilter)
		Verbose.Printf("Filtering for %s\n", g_extFilter)
	}

	// -cpu=n
	var NUM_CORES int = runtime.NumCPU()
	Verbose.Printf("CPUs from command line = %d\n", flagCPU)
	Verbose.Printf("NumCPU(%d)\n", NUM_CORES)
	Verbose.Printf("GOMAXPROCS(%q)\n", os.Getenv("GOMAXPROCS"))
	if flagCPU != 0 { // it was set, so force to reasonable value
		nCPU = flagCPU
		if flagCPU >= NUM_CORES {
			nCPU = NUM_CORES
		}
		if flagCPU < 0 {
			nCPU = 1
		}
	} else { // default to MAX
		nCPU = NUM_CORES
	}
	Verbose.Printf("setting GOMAXPROCS to %d (nCPU)\n", nCPU)
	runtime.GOMAXPROCS(nCPU)
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Printf("Nothing to do - No arguments in command line\n")
		usage()
		return
	}
	if g_verboseFlag {
		Verbose = true
	}
	flagSetup()

	// BUG(mdr): TODO? flag for relative path or flag for abs path?
	pathName, err := filepath.Abs(flag.Arg(0))
	if err != nil {
		log.Fatalf("cant get absolute path for %s\n", flag.Arg(0))
	}

	lo := make(chan digestType, 5)
	//fmt.Printf("before lo %d gort running\n",runtime.NumGoroutine())
	go lineOut(lo)

	Verbose.Printf("Checking paths in %s\n", pathName)
	dirInfo, err := os.Stat(pathName)
	if err != nil {
		log.Fatalf("cant stat the directory %s\n", pathName)
	}
	dMode := dirInfo.Mode()
	if dMode.IsDir() == false {
		log.Fatalf("Path %s must be a directory (but isn't)\n", pathName)
	} else {
		Verbose.Printf("%s is a directory, walking starts now\n", pathName)
	}
	filepath.Walk(pathName, CheckPath) // builds g_argList
	var filesProcessed int64 = 0
	var bytesProcessed int64 = 0
	fmt.Fprintf(os.Stderr, "# nCPU = %d\n", nCPU)
	throttle := make(chan int, nCPU)
	startTime := time.Now()
	for _, fname := range g_argList {
		//fmt.Printf("Goroutines active = %d\n", runtime.NumGoroutine())
		throttle <- 1
		loop.Add(1)
		go func(fullpath string, accel chan int) {
			defer loop.Done()
			var tmp digestType
			tmp.pathname = fullpath
			stats, err := os.Stat(fullpath)
			if err != nil {
				log.Fatalf("Can't get fileinfo for %s\n", fullpath)
			}
			// check time for sanity (date < now()
			tmp.fileDate = stats.ModTime()
			if tmp.fileDate.After(startTime) {
				fmt.Printf("# bad date %s for %s\n", tmp.fileDate.String(), fullpath)
				// BUG(mdr): TODO? save bad dates in a list for appending
				g_badDateCt++
			}
			if g_noSHAFlag {
				// do nothing
			} else {
				tmp.dig256, err = mdr.FileSHA256(fullpath)
				if err != nil {
					log.Fatalf("SHA256 failed on %s\n", fullpath)
				}
			}
			tmp.fileLength = stats.Size()
			g_tmMutex.Lock()
			bytesProcessed += tmp.fileLength
			filesProcessed++
			g_tmMutex.Unlock()
			lo <- tmp
			<-accel // free a core
		}(fname, throttle)
	}
	loop.Wait()
	var doneRec digestType
	doneRec.pathname = ""
	lo <- doneRec
	outPut.Wait()
	//time.Sleep(1 * time.Second) // not necessary
	// wrapup
	elapsedTime := time.Now().Sub(startTime)
	elapsedSeconds := elapsedTime.Seconds()
	fmt.Printf("# %s Rundate=%s\n", G_version, startTime.String())
	fmt.Printf("# Processed %s files with %s bytes in %s for %.2g bytes/sec\n",
		mdr.CommaFmtInt64(filesProcessed), mdr.CommaFmtInt64(bytesProcessed), mdr.HumanTime(elapsedTime), float32(bytesProcessed)/float32(elapsedSeconds))
	fmt.Printf("# nCPU[%d]  BadDates[%d]\n", nCPU, g_badDateCt)
}
