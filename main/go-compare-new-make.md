# go compare new | make

        make只适用于chan map slice;
        make可以指定len cap
        make返回类型本身

## new

```go
// 分配内存后, 返回指针
func new(Type) *Type

// new不常用, 一般用短语句声明 or 结构体字面量替代
i:=0
u:=user{}
```

## make

```go
// 分配内存后, 返回类型 (make适用于chan, map, slice都是引用类型, 所以返回类型本身)
func make(t Type, size ...IntegerType) Type

// make无可替代, 初始化chan, map, slice时必须用到
```

## 相同点

- 都是分配内存

## 不同点

- new适用于任何类型, make适用于chan map slice
- new返回指针, make返回类型值本身(因为都是引用类型)
- make可以指定len cap

## ref

- <https://www.flysnow.org/2017/10/23/go-new-vs-make.html>