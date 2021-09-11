# go bufio

## code

- [bufio](../script/go/package/bufio_test.go)

## bufio read()

            len(p)<len(buf)         len(buf)==0, 读file, 填满buf
    program <------------------ buf <------------------ file
            len(p)>len(buf)
    program <------------------------------------------ file

> `len(p)`读取的内容的长度

## bufio write()

            len(p)<len(buf)         len(buf)>0 && len(p)>len(buf), 填满buf, 写入file, 清空buf
    program ------------------> buf ------------------> file
            len(buf)==0 && len(p)>len(buf) 
    program ------------------------------------------> file

> `len(p)`写入的内容的长度