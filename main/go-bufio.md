# go bufio

## code

- [bufio](src/go/package/bufio_test.go)

## bufio read()

            len(p)<len(buf)         len(buf)==0, 读file, 填满buf
    program <------------------ buf <------------------ file
            len(p)>len(buf)
    program <------------------------------------------ file

> `len(p)`读取的内容的长度

### read示例

```bash
# 从硬盘读取10个字符, 每次读取1个, bufio利用缓冲, 实际只从硬盘读取3次
io.Reader --> buffer --> consumer

    disk-io       mem
abcd -----> abcd -----> a
            abcd -----> b
            abcd -----> c
            abcd -----> d
efgh -----> efgh -----> e
            efgh -----> f
            efgh -----> g
            efgh -----> h
ijkl -----> ijkl -----> i
            ijkl -----> j
```


## bufio write()

            len(p)<len(buf)         len(buf)>0 && len(p)>len(buf), 填满buf, 写入file, 清空buf
    program ------------------> buf ------------------> file
            len(buf)==0 && len(p)>len(buf) 
    program ------------------------------------------> file

> `len(p)`写入的内容的长度

### write示例

```bash
# 频繁写入, 利用bufio的缓冲, 攒一波写入(flush)一次, 减少disk-io
producer --> buffer --> io.Writer (destination )

       disk-io           mem
   a    ----->    a
   b    ----->    ab
   c    ----->    abc
   d    ----->    abcd
   e    ----->    e      ----->   abcd
   f    ----->    ef
   g    ----->    efg
   h    ----->    efgh
   i    ----->    i      ----->   abcdefgh
```