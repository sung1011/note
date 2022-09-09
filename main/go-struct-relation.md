# go struct relation

## 包含

```go
type Animal struct {
    name   string
    weight int
    father *Animal // 同类扩展; Animal不可以, *Animal可以
}

bm_horse := &Animal{ // 马 和 ta的爸爸
    name: "baima",
    weight: 60,
    father: &Animal{
        name: "hongma",
        weight: 80,
        father: nil,
    },
}
```

> T类型中包含T类型字段 不合法; 但可以包含*T, []T, map[string]T

> T1包含T2, T2包含T1 不合法

## 内嵌

```go
type Animal struct {
    Name   string
    Weight int
}

type Horse struct {
    *Animal // 内嵌; 可以理解成嵌到Horse下面, 因为Animal访问不到Horse
    Speak string
    int //匿名字段, 一般表达继承
}

func (a *Animal) hello() {
    fmt.Println(a.Name)
    fmt.Println(a.Weight)
    //fmt.Println(a.Speak) // animal访问不到horse的字段,  即便是horse实例
    // fmt.Println(a.int) // animal访问不到horse的匿名字段
}

func (h *Horse) hello() { // 重载 覆盖 Animal.hello()
    fmt.Println(h.Name)
    fmt.Println(h.Weight)
    fmt.Println(h.Speak)
    fmt.Println(h.int)
}

func main() {
    bm_horse := &Horse{
        Animal: &Animal{ // 结构体包含`嵌入字段`, 类型名充当了嵌入字段的字段名
            Name:   "baima",
            Weight: 60,
        },
        Speak: "neigh",
        int: 1234,
    }
    bm_horse.hello() // horse.hello()
}
```

## 冲突

```go
type A struct {
	a int
	b int
}

type B struct {
	b float32
	c string
	d string
}

type C struct {
	A
	B
	a string
	c string
}

func main() {
	var c C
	c.a = "a"        // C的a高于A的a
	c.A.a = 123      // A的a

	// C.b = 1.23    // 冲突, 不确定A.b 或 B.b, 需要指明; error: C.b undefined (ambiguous selector)
}

```

## 实例

```go
// *file是嵌入指针, 外层的File保障了*file不会被复制(没有获取*file的方法)
type File struct {
    *file
}
```

```go
// 无法访问另一个包公开类型的未公开字段
// 可以访问另一个包公开类型的未公开字段(值是类型)的公开字段 ???
// TODO
```
