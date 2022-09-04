# go type

## 值类型

    变量直接存储值, 内存通常在栈中分配, 栈在函数调用完会被释放

- ptr  
- [array](go-array.md)  
- [struct](go-struct.md)  
- bool  
- [int](go-int.md)
- [float](go-float.md)
- [string](go-string.md)  
- unsafe.Pointer

## 引用类型

    变量存储的是一个地址, 这个地址存储最终的值. 内存通常在堆上分配, 通过GC回收.

- [slice](go-slice.md)  
- [map](go-map.md)  
- [chan](go-chan.md)  
- func
- [interface](go-interface.md)

## 自定义类型 与 类型别名

- [custom](go-type-custom.md)

## size

| type        | size (byte) | value range                                |
| ----------- | ----------- | ------------------------------------------ |
| bool        | 1           | true/false                                 |
| ptr         | 8           |                                            |
| channel     | 8           | ptr                                        |
| map         | 8           | ptr                                        |
| func        | 8           | ptr                                        |
| slice       | 24          | ptr + len + cap                            |
| array       | ~           | T * len                                    |
| struct      | ~           |                                            |
| string      | 16          | ptr + len                                 |
| uint8       | 1           | 0~255                                      |
| uint16      | 2           | 0~65535                                    |
| uint32/rune | 4           | 0~4294967295                               |
| int64/int   | 8           | 0~18446744073709551615                     |
| int8        | 1           | -128~127                                   |
| int16       | 2           | -32768 ~ 32767                             |
| int32/rune  | 4           | -2147483648 ~ 2147483647                   |
| uint64/int  | 8           | -9223372036854775808 ~ 9223372036854775807 |
| float32     | 4           | -2147483648 ~ 2147483647                   |
| float64     | 8           | -9223372036854775808 ~ 9223372036854775807 |

> `unsafe.Sizeof(T)` size

> `int` 4byte in 32bit; 8byte in 64bit

> `timestamp` 如 1604556795 # int32: 2^32=4294967296 能容纳 10位数
