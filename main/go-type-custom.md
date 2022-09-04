# go 自定义类型

## 创建新类型

### 声明

    基于基础类型, 声明一个新类型

```go
    type myInt int
```

### 转化

    类型不同, 必须显式转化

```go
    type myInt int // int 是 myInt的底层类型(underlying type)

    var i int = 123
	// var mi myInt = i         error: cannot use i (type int) as type myInt in assignment
	var mi myInt = myInt(i)
```

## 类型别名 Type alias

    类型相同, 可以互转, 只是起了个别名
    一般用于 别名表达, 渐进式重构

```go
    type myInt = int

    var i int = 123
    var mi myInt = i
```