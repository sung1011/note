# go float

## type

| type    | size (byte) | value range                                |
| ------- | ----------- | ------------------------------------------ |
| float32 | 4           | -2147483648 ~ 2147483647                   |
| float64 | 8           | -9223372036854775808 ~ 9223372036854775807 |

> [size](go-type-size.md)

## 溢出

todo

## 格式化

```go
    3.1415
    .333         // 0.333
    81.30
    12.          // 12
    12.3e3       // 12.3 * 10^3 = 12300
    12.3E-3      // 12.3 * 10^-3 = 0.0123
    0x2.p10      // 2.0 * 2^10 = 2048; 十六进制科学计数法
    0x1.Fp+0     // 1.9375 * 2^0 = 1.937500
    fmt.Printf() // %e 十进制的科学计数法; %x 十六进制的科学计数法
```

## 运算

    使用第三方包 github.com/shopspring/decimal

```go
	var v float64 = 67.6
	fmt.Println(v * 100)    // 6759.999999999999

    f1 := decimal.NewFromFloat(v)
    f2 := decimal.NewFromFloat(100)
    fmt.Println(f1.Mul(f2).IntPart()) // 6760
```

## 比较

    使用第三方包 github.com/shopspring/decimal
    转化为string进行比较

```go
	var f1 float32 = 16777216.
	var f2 float32 = 16777217.
	fmt.Println(f1 == f2)       // true

	v1 := decimal.NewFromFloat32(f1)
	v2 := decimal.NewFromFloat32(f2)
	fmt.Println(v1 == v2)       // false
```