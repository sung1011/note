# go type

## 值类型

变量直接存储值，内存通常在栈中分配，栈在函数调用完会被释放

- bool  
- [int](go-int.md)
- float  
- string  
- ptr  
- [array](go-array.md)  
- [struct](go-struct.md)  
- unsafe.Pointer

## 引用类型

变量存储的是一个地址，这个地址存储最终的值。内存通常在堆上分配，通过GC回收。

- [slice](go-slice.md)  
- [map](go-map.md)  
- [chan](go-chan.md)  
- [interface](go-interface.md)
- func
