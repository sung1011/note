# go compare

## os, bufio, ioutil

- `os` 调用系统函数
- `bufio` 带缓冲区; 最优选
- `ioutil` 1次IO, 分配大量内存; 易用, 小文件优选

### read

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

> gb级别 ioutil需要一次性开辟GB级的内存; os每次io都调用系统read(); bufio则是优先读缓冲区, 所以更快

#### bufio 读取一行 ReadBytes, ReadString, ReadSlice, ReadLine

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

## gzip, zip, tar

### zip

```txt
more popular on Windows
archiving and compression
use DEFLATE compression algorithm (same gzip)
```

### gzip

```txt
more popular on Linux/Unix
faster than ZIP
more save space than ZIP
just compression
use DEFLATE compression algorithm (same zip)
```

### tar

```txt
archive(package) the files and dirs
```

