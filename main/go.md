# go  
  
## type  

### 值类型

变量直接存储值，内存通常在栈中分配，栈在函数调用完会被释放

- bool  
- int  
- float  
- string  
- [struct](go-struct.md)  
- [array](go-array.md)  
- ptr  
- [interface](go-interface.md)
- func
- unsafe.Pointer

### 引用类型

变量存储的是一个地址，这个地址存储最终的值。内存通常在堆上分配，通过GC回收。

- [slice](go-slice.md)  
- [map](go-map.md)  
- chan  

## [标准库](go-package.md)

## [并发](go-concurrent.md)

## [test](go-test.md)

## GC
  
## ref

[标准库](https://studygolang.com/pkgdoc)  
