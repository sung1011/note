# complexity

## 时间复杂度

常数 O(1)  
对数 O(log(n))  
线性 O(n)  
线性对数 O(nlog(n))  
平方 O(n², n^2)  
立方 O(n³, n^3)  
指数 O(2^n)  
阶乘 O(n!)

## 空间复杂度

## 实战

### O(log(n))

```go
    n := 100
    for i := 1; i < n; i = i * 2 {
        fmt.Println(i) //[7] 1 2 4 8 16 32 64
    }
```

### O(n^2)

```go
    n := 3
    for i := 1; i <= n; i++ {
        for j := 1; j <= n; j++ {
            fmt.Println(i, "~", j) //[9] 1~1 1~2 1~3 2~1 2~2 2~3 3~1 3~2 3~3
        }
    }
```

### O(2^n)

```go
    var n float64 = 4
    var i float64
    for i = 1; i <= math.Pow(2, n); i++ {
        fmt.Println(i) //[16] 1...16
    }
```

### O(n!)

```go
    for i := 1; i < factorial(n); i++ {
        fmt.Println(i)
    }
```
