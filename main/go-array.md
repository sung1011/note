# go 数组

## 数据结构

- `len`
- `type`

## 创建 初始化 访问

```go
var arr [5]int;
// var arr [...]int; error: use of [...] array outside of array literal
arr := [5]int;
arr := [5]int{10, 20, 30, 40, 50};
arr := [...]int{10, 20, 30, 40, 50};
arr := [...]string{ 99: "foo" } // 100个元素
arr := [...]*int{new(int), new(int)}
```

## 赋值

    相同数据结构(即len + type相同), 可以赋值

```go
a1 := [3]int{1,2,3};
a2 := [3]int{8,9}   // 只能是[3]int, 如是[2]int, [4]int都会报错
a2[0] = 888         // 修改元素的值; 888
a1 = a2             // 覆盖赋值; 888, 9, 0
```

## [是否==](go-type-compare.md#array)

## 多维数组

```go
arr := [4][2]int{
    {1,2}, {2,2}, {1,1}, {2,1},
}
```

## 数组参数

      &arr 传递指针 8字节
      arr 传递值 T*len 可能巨大
