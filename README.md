
<center>
## ls256
</center>

License details are at the end of this document. 
This document is (c) 2014 David Rook.

Comments can be sent to <hotei1352@gmail.com>

### Description

This [package][1] compiles with go 1.2 (and probably some older versions as well).
The [interface docs][2] are at godoc.org.

	list256 dir
	typical output line:     6003568 | 9df360b8ec6a58f6cb410303c7794d98526cbfe2b11be7a34754515d0fcb21bb | 2013-07-03:09_03_57 | /home/mdr/Desktop/MYGO/src/vwar3/vwar3

	First field is size.  Second is SHA256 of the file. Third is date.  Fourth is full pathname.
	
	By convention the output is saved to filename.256 using 'tee'.
	
### Why it's useful

When used with other programs it allows us to find (and potentially eliminate) duplicate files and directories.  In other contexts it
can be used to detect bit-rot or tampering of files.  As of the time of writing (2014) SHA256 is considered secure - to the extent that it
is believed by experts in the field to be "extremely difficult" to construct a file that has the same SHA256 and size but different contents
than a given file.

Comments (signaled by # in first column) are included at the end of processing which indicate
when the run was made, how many files and bytes were processed, how long the run took,
how many CPUs were in use, and how many dates were bad. A bad date is counted when it sees a 
file with a date that's in the "future" as of the time the program was started.

### Companion programs
* onecopy256 - finds duplicate files in a single directory tree
* two256 - finds duplicate files between two trees
* undupeTree256 - finds duplicate directory trees
* mdr package
	* funcs for loading *.256 output as list or map (indexed by SHA256 or Pathname)  
	* func for splitting *.256 lines into FileRec struct
	
### Limitations 
	
* Program gathers all filenames as an argument list before procesing.  This can exhaust memory if too many files.  On the other hand it
makes it less likely that stat() will fail when the arguments are evaluated.

---

### Usage

```
ls256 [OPTIONS] path

OPTIONS:
	-verbose               more intermediate messages
	-nosha                 dont compute SHA256  (still does size date and name)
	-cpu=n                 limit operation to n CPUs (default is all available)
	-ext=".ext"            limit to files with this extension (default is all)
	-links                 include directory link count in output

```

### Installation

```
go get github.com/hotei/ls256
```


---

### Resources

* [Source for package] [1]

---

### Change Log

* 2014-04-01 - dirlink as flag, default false
* 2014-03-22 - compute  dirlink count during path walk
* 2013-03-19 - started

[1]: http://github.com/hotei/ls256 "github.com/hotei/ls256"
[2]: http://godoc.org/github.com/hotei/ls256 "godoc.org"

### License

The 'ls256' go package and demo programs are distributed under the Simplified BSD License:

> Copyright (c) 2014 David Rook. All rights reserved.
> 
> Redistribution and use in source and binary forms, with or without modification, are
> permitted provided that the following conditions are met:
> 
>    1. Redistributions of source code must retain the above copyright notice, this list of
>       conditions and the following disclaimer.
> 
>    2. Redistributions in binary form must reproduce the above copyright notice, this list
>       of conditions and the following disclaimer in the documentation and/or other materials
>       provided with the distribution.
> 
> THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDER ``AS IS'' AND ANY EXPRESS OR IMPLIED
> WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND
> FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> OR
> CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
> CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
> SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
> ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
> NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
> ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// EOF README-ls256.md  (this is a markdown document and tested OK with blackfriday)
