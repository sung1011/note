# go const

const NAME [TYPE] = VALUE  
TYPE只能是 bool, string, number(int/float/complex)

```go
const Pi float32 = 3.14159

const c = 3+2         

// const d = getNumber()  // error; 类型错误

const Ln2= 0.693147180559945309417232121458\    // 可以用\换行
            176568075500134360255254120680009

const beef, two, c = "meat", 2, "veg"

const ( // 枚举
    Unknown = 0
    Female = 1
    Male = 2
)

const ( // iota
	a    = iota           // 0: 0
	b                     // 1: 1
	g    = iota + 5       // 2: 2+5 = 7
	h                     // 3: 8; 延续自增值
	i    = iota           // 4: 4; iota是索引值
	j, k = iota, iota + 2 // 5: 5, 7; 平级不变
	l    = iota           // 6: 6
)
const (
	x = iota // 0: 0 新代码块 重置
)
```
