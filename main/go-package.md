# go package

## [import](go-import.md)

## compare

### os, bufio, ioutil

- `os` 调用系统函数
- `bufio` 带缓冲区; 最优选
- `ioutil` 1次IO, 分配大量内存; 易用, 小文件优选

#### read

| type   | size | cost-time |
| ------ | ---- | --------- |
| os     | 2k   | 363µs     |
| bufio  | 2k   | 50µs      |
| ioutil | 2k   | 34µs      |

> kb级别 差别不大, 因为io最少, ioutil最快

| type   | size | cost-time |
| ------ | ---- | --------- |
| os     | 21M  | 16ms      |
| bufio  | 21M  | 8ms       |
| ioutil | 21M  | 34ms      |

> mb级别 差别不大, bufio的优势开始显现

| type   | size | cost-time |
| ------ | ---- | --------- |
| os     | 1G   | 385ms     |
| bufio  | 1G   | 367ms     |
| ioutil | 1G   | 4s        |

> gb级别 ioutil需要一次性开辟gb的内存; os每次io都调用系统read(); bufio则是优先读缓冲区, 所以更快

##### bufio 读取一行 ReadBytes, ReadString, ReadSlice, ReadLine

- `ReadBytes` 循环make, append, copy比较耗时; 返回copy
- `ReadString` 调用`ReadBytes()`; 返回copy
- `ReadSlice` 切片式读取; 返回buffer
- `ReadLine` 调用`ReadSlice()`, 只是对换行符做了优化; 返回buffer; 最优选.

| type       | size | cost-time |
| ---------- | ---- | --------- |
| ReadBytes  | 2k   | 512µs     |
| ReadString | 2k   | 46µs      |
| ReadSlice  | 2k   | 23µs      |
| ReadLine   | 2k   | 18µs      |

> kb级别 相差不大

| type       | size | cost-time |
| ---------- | ---- | --------- |
| ReadBytes  | 21M  | 31ms      |
| ReadString | 21M  | 39ms      |
| ReadSlice  | 21M  | 32ms      |
| ReadLine   | 21M  | 18ms      |

> mb级别 相差不大, ReadLine较快

| type       | size | cost-time |
| ---------- | ---- | --------- |
| ReadBytes  | 1G   | 1s        |
| ReadString | 1G   | 1s        |
| ReadSlice  | 1G   | 692ms     |
| ReadLine   | 1G   | 580ms     |

> gb级别 ReadLine, ReadSlice较快

### gzip, zip, tar

## zip

```txt
more popular on Windows
archiving and compression
use DEFLATE compression algorithm (same gzip)
```

## gzip

```txt
more popular on Linux/Unix
faster than ZIP
more save space than ZIP
just compression
use DEFLATE compression algorithm (same zip)
```

## tar

```txt
archive(package) the files and dirs
```


## pkg

### code

#### [file](../script/go/file/main_test.go)

### describe

#### [archive](go-archive.md)

- tar
- zip

#### [bufio](go-bufio.md)

#### TODO [builtin](_)

#### TODO [bytes](_)

#### TODO [compress](_)

- bzip2
- flate
- gzip
- lzw
- zlib

#### [context](go-context.md)

#### TODO [container](_)

- heap
- list
- ring

#### [crypto](go-crypto.md)

- aes
- cipher
- des
- dsa
- ecdsa
- ed25519
- elliptic
- hmac
- md5
- rand
- rc4
- rsa
- sha1
- sha256
- sha512
- subtle
- tls
- x509
- x509/pkix

#### TODO [database](_)

- sql
- sql/driver

#### TODO [debug](_)

- dwarf
- elf
- gosym
- macho
- pe
- plan9obj

#### TODO [embed](_)

#### [encoding](go-encoding.md)

- ascii85
- asn1
- base32
- base64
- binary
- csv
- gob
- hex
- json
- pem
- xml

#### [errors](go-errors.md)

#### TODO [expvar](_)

#### TODO [flag](_)

#### [fmt](go-fmt.md)

#### TODO [go](_)

- ast
- build
- build/constraint
- constant
- doc
- format
- importer
- parser
- printer
- scanner
- token
- types

#### TODO [hash](_)

- adler32
- crc32
- crc64
- fnv
- maphash

#### TODO [image](_)

- color
- color/palette
- draw
- gif
- jpeg
- png

#### TODO [index](_)

- suffixarray

#### TODO [io](_)

- fs
- ioutil

#### TODO [log](_)

- syslog

#### TODO [math](_)

- big
- bits
- cmplx
- rand

#### TODO [mime](_)

- multipart
- quotedprintable

#### TODO [net](_)

- http
- http/cgi
- http/cookiejar
- http/fcgi
- http/httptest
- http/httptrace
- http/httputil
- http/pprof
- mail
- rpc
- rpc/jsonrpc
- smtp
- textproto
- url

#### TODO [os](_)

- exec
- signal
- user

#### TODO [path](_)

- filepath

#### TODO [plugin](_)

#### TODO [reflect](go-reflect.md)

#### TODO [regexp](_)

- syntax

#### TODO [runtime](_)

- cgo
- debug
- metrics
- pprof
- race
- trace

#### TODO [sort](_)

#### TODO [strconv](_)

#### TODO [strings](_)

#### TODO [sync](_)

- atomic

#### TODO [syscall](_)

- js

#### TODO [testing](_)

- fstest
- iotest
- quick

#### TODO [text](_)

- scanner
- tabwriter
- template
- template/parse

#### TODO [time](_)

- tzdata

#### TODO [unicode](_)

- utf16
- utf8

#### TODO [unsafe](_)

#### TODO [internal](_)

- bytealg
- cfg
- cpu
- execabs
- fmtsort
- goroot
- goversion
- lazyregexp
- lazytemplate
- nettrace
- obscuretestdata
- oserror
- poll
- profile
- race
- reflectlite
- singleflight
- syscall/execenv
- syscall/unix
- syscall/windows
- syscall/windows/registry
- syscall/windows/sysdll
- sysinfo
- testenv
- testlog
- trace
- unsafeheader
- xcoff

## ref

<https://pkg.go.dev/std>