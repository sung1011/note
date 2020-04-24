# go 切片

## 数据结构

slice: `ptr`指向底层array第一个元素地址, `len`, `cap`  
array: {0:0, 1:0, 2:0, 3:0 ...}  

## 创建 初始化 访问

```go
sl := make([]int, 5)
sl := make([]int, 3, 5)
sl := []int{30, 50, 20}
```

## 数组 与 切片

```go
a := [...]int{1, 2, 3, 4, 5}
sl := a[1:3]
fmt.Println(sl, len(sl), cap(sl)) //[2 3] len = 2 cap = 4
```

## nil切片 & 空切片

nil 切片: `var nil_sl []int`  ptr为nil, len 0, cap 0  
空 切片: `empty_sl := make([]int, 0)` ptr指向底层的空array, len 0, cap 0

## 拷贝 copy

```go fake
//go doc builtin copy
n := copy({1,2,3}, {111, 222, 333, 444}) //sl1 = {111, 222, 333}; n = 3
n := copy({1,2,3}, {111, 222}) //sl1 = {111, 222, 3}; n = 2
```

## append & 扩容

向切片增加元素。  
cap不足时，新建底层数组并扩容cap。  
cap足够时，会改变底层数组元素值。  

```go
{1, 2, 3, 4} := append({1, 2}, 3, 4)
{1, 2, 3, 4} := append({1, 2}, {3, 4}...)
```

## 作为参数

slice{ptr, len, cap}的副本, 值传递给函数作为参数

## 内存GC

问题: 当切片对应的底层数组很大，而GC不会回收正在被引用的对象，造成内存浪费。  
解决: 当函数的返回值是指向底层数组的数据结构(如slice)，应在函数内copy slice到新slice并返回新slice。
效果: 函数退出时老slice对应的较大的底层数组会被回收，保存在内存的是新的小slice。

```go
func foo() {
old_sl := ... //切片返回值
new_sl := make([]T, len(old_sl)) //新的底层数组
copy(new_sl, old_sl) //新切片
return new_sl //新切片返回值
}
```

## ref

[slice底层实现](https://blog.csdn.net/lengyuezuixue/article/details/81197691)
