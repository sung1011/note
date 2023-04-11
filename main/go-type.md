# 1. go type

- [1. go type](#1-go-type)
  - [1.1. 类型](#11-类型)
    - [1.1.1. 值类型](#111-值类型)
    - [1.1.2. 引用类型](#112-引用类型)
    - [1.1.3. 自定义类型 与 类型别名](#113-自定义类型-与-类型别名)
  - [1.2. 是否可==](#12-是否可)
  - [1.3. size](#13-size)
  - [1.4. ref](#14-ref)

## 1.1. 类型

    零值不同
        引用 nil (即 可以用==比较nil)
        值 对应的零值

    初始化方式不同
        引用 make(), new()
        值 不必初始化

### 1.1.1. 值类型

    变量直接存储值, 内存通常在栈中分配, 栈在函数调用完会被释放

- [array](go-array.md)  `不可== (不同len)`
- [struct](go-struct.md)  `不可== (含不可比较类型)`
- [int](go-int.md)
- [float](go-float.md)
- [string](go-string.md)  
- bool  

### 1.1.2. 引用类型

    变量存储的是一个地址, 这个地址存储最终的值. 内存通常在堆上分配, 通过GC回收.

- [slice](go-slice.md)  `不可==`
- [map](go-map.md)  `不可==`
- [func](go-func.md)  `不可==`
- [chan](go-chan.md)
- [interface](go-interface.md)
- [ptr](go-ptr.md)

### 1.1.3. 自定义类型 与 类型别名

- [custom](go-type-custom.md)

## 1.2. [是否可==](go-type-compare.md)

## 1.3. [size](go-type-size.md)

## 1.4. ref

- 值类型 指针类型 引用类型 值传递 指针传递 <https://studygolang.com/articles/29675>