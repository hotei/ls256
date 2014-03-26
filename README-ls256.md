
<center>
ls256
========
</center>

License details are at the end of this document. 
This document is (c) 2014 David Rook.

Comments can be sent to <hotei1352@gmail.com>

Description
-----------
This [package][1] Compiles with go 1.2 (and probably some older versions as well).
The [interface docs][2] are at godoc.org.

	list256 dir
	typical output:     6003568 | 9df360b8ec6a58f6cb410303c7794d98526cbfe2b11be7a34754515d0fcb21bb | 2013-07-03:09_03_57 | /home/mdr/Desktop/MYGO/src/vwar3/vwar3

	limitation - gathers filenames as an argument list before procesing.  Can exhaust memory if too many files
	benefit - potentially allows use of progress bar since we know how many items are to be processed
	  however, we're not currently using progress bar


---

Installation
------------

```
go get github.com/hotei/ls256
```


---

Resources
---------
* [Source for package] [1]

---

Journal
-------
* 2014-03-22 - compute  dirlink count during path walk
* 2013-03-19 - started

[1]: http://github.com/hotei/ls256 "github.com/hotei/ls256"
[2]: http://godoc.org/github.com/hotei/ls256 "godoc.org"
License
-------
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
