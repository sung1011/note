# go 结构体

## 数据结构

        TODO

## 创建 初始化 访问

```go
type Student struct {// 结构体
    Name string
    Age int
    Scores map[string]int
}
// 声明
// 零值可用, 不必初始化, 方便.如: var b bytes.Buffer, var mu sync.Mutex
var s Student 
s.Name = "Sam"
s.Age = 18
// 初始化
sp1 := new(Student) 
// 初始化 并赋值; 不推荐赋值时 按顺序匿名赋值
sp2 := &Student{Name: "Sam", Age: 18, Scores: make(map[string]int)}
sp2.Scores = map[string]int{"math": 98, "chemistry": 80}
// 初始化 并顺序初始值
sp3 := &Student{"Sam", 18} 
```

## 内存布局 连续

    结构体和它所包含的数据在内存中是以连续块的形式存在的, 即使结构体中嵌套有其他的结构体

![img](res/go-struct-assign.jpg)

![img](res/go-struct-mem.jpg)

## 参数传递

    一般传递其指针 &struct (8byte)

## 方法

```go
type User struct {
	name   string
	School *SchoolInfo // 当期望 不随User变化而变化, 使用指针
	_      int8        // 匿名字段
}
type SchoolInfo struct {
	Addr string
}

// 指针接收者 修改类型实例本身 (不论调用者是u或&u)
func (u *User) changeName() {
	u.name = "xx"
}

// 值接收者 修改类型实例的副本 (不论调用者是u或&u)
// func (u User) changeName() User {
// 	u.name = "xx"
// 	return u
// }

// 获取内嵌字段
func (u User) getSchoolAddr() string {
	return u.School.Addr
}

func main() {
	u := &User{ // 初始化
		name:   "sun",
		School: &SchoolInfo{Addr: "xxx"}, // 初始化内嵌结构
	}
	fmt.Println("", u)
}
```

> 关注Type是值类型还是引用类型, 引用类型用值接收者即可, 值类型区分区分是修改实例本身or生成副本

> 如果一个工厂函数返回的是指针, 其方法都应当使用指针接收者 (即使该方法不修改实例本身); 例外: 需要类型值符合某个接口时

> func(recv) 是函数; recv.Method() 是方法

> T类型中包含T类型字段 不合法; 但可以包含*T, []T, map[string]T

> T1包含T2, T2包含T1 不合法

## 基于基础类型的新类型

```go
type Duration int64

var i int64
i = 100
// j := i // 最后一行赋值编译会报错; 需要显式转换类型 `j := Duration(i)`
j := Duration(i) // 可以赋值是因为他们具有相同的底层类型(underlying type)int64
var dur Duration
dur = j
```

## 构造函数

    作为工厂方法用于构造实例, 返回实例的指针
    特别是可以对未导出的字段赋值

```go
// func NewT(field1, field2, ...) *T
func NewStudent(name string, age int) *Student {
    return &Student{
        name: name,
        age: age,
    }
}
```

## 包含

```go
type Animal struct {
    name   string
    weight int
    father *Animal // 同类扩展
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

## 组合 & 重载 & 内嵌

```go
type Animal struct {
    Name   string
    Weight int
}

type Horse struct {
    *Animal //组合, 内嵌
    Speak string
    int //匿名字段
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

## 标签tag

```go
type Server struct {
    Name string `tag:"名字,域名"`
    IP   string `tag:"ip地址" size:"30"`
}

func main() {
    t := reflect.TypeOf(Server{})
    tn, _ := t.FieldByName("IP")
    fmt.Println(tn.Tag.Get("size")) // 30
}
```

## [嵌套接口](go-interface.md#嵌入结构体)

## 未公开类型的公开字段

```go
    TODO
```

## 空结构体 struct{}{}

    不占用内存
    地址不变

```go
    type Empty struct{}
    var s Empty
    fmt.Println(unsafe.Sizeof(s)) // 0

    var ch := make(chan Empty)
    ch <-Empty  // 不占用内存, 只作为一种事件信息传递
```

## 内存对齐

    每个字段占用8byte, 字段顺序影响其内存空间占用

```go
type T1 struct {
	b byte   // 1
	i int64  // 8
	u uint16 // 2
}
type T2 struct {
	b byte   // 1
	u uint16 // 2
	i int64  // 8
}

func main() {
	var t1 T1 // (1+`7`) + 8 + 2 = 24; 7为填充
	var t2 T2 // (1+2+`5`) + 8 = 16; 5为填充
	fmt.Println(unsafe.Sizeof(t1), unsafe.Sizeof(t2))
}
```

## 链表 与 双向链表

![img](res/go-struct-list.jpg)

```go
type Node struct { // 链表
    data    float64
    su      *Node // 下一节点指针
}

type Node struct { // 双向链表
    data float64
    su *Node // 下一节点指针
    pr *Node // 上一节点指针
}
```

## 二叉树

![img](res/go-struct-btree.jpg)

```go
type Tree strcut {
    le      *Tree
    data    float64
    ri      *Tree
}
```

## ref

- <https://www.cnblogs.com/myuniverse/p/11595043.html>
