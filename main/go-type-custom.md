# go 自定义类型

## 创建新类型

### 声明

    基于基础类型, 声明一个新类型
    新类型不是继承底层类型, 所以不包含底层类型的方法

```go
    type myInt int
```

### 转化

    类型不同, 必须显式转化

```go
    type Duration int64 // int64 是 Duration 的底层类型(underlying type)

	var i int64 = 100
	// j := i // error: cannot use j (type int64) as type Duration in assignment
	j := Duration(i) // 可以转换, 他们具有相同的底层类型(underlying type)int64
	var dur Duration = j
	fmt.Println("", dur)
```

## 类型别名 Type alias

    类型相同, 可以互转, 只是起了个别名
    一般用于 别名表达, 渐进式重构

```go
    type myInt = int

    var i int = 123
    var mi myInt = i
```