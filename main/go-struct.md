# go 结构体

## 数据结构

        TODO

## 创建 初始化 访问

```go
type Student struct {// 结构体
    Name string
    Age int
}

var s Student // 声明
s.Name = "Sam" // 赋值
s.Age = 18

sp := new(Student) // 初始化

sp := &Student{ // 初始化 与new等效
    Name    :"Sam",
    Age     :18,
}

sp := &Student{"Sam", 18} // 初始化 with顺序初始值
```

![img](res/go-struct-assign.jpg)

## 内存布局 连续

    结构体和它所包含的数据在内存中是以连续块的形式存在的, 即使结构体中嵌套有其他的结构体

![img](res/go-struct-mem.jpg)

## 参数传递

    一般传递其指针 &struct (8byte)

## 方法

```go
type User {
    name string
    school *schoolInfo // 当期望 不随User变化而变化, 使用指针
}
// 指针接收者 修改类型实例本身 (不论调用者是u或&u)
func (u *User) changeName() {
    u.name = "xx"
}
// 值接收者 修改类型实例的副本 (不论调用者是u或&u)
func (u User) changeName() User {
    u.name = "xx"
    return u
}
```

> 关注Type是值类型还是引用类型, 引用类型用值接收者即可, 值类型区分区分是修改实例本身or生成副本

> 如果一个工厂函数返回的是指针, 其方法都应当使用指针接收者 (即使该方法不修改实例本身); 例外: 让类型值符合某个接口时.

> func(recv) 叫函数; recv.Method() 叫方法

## 基于基础类型的新类型

```go
type Duration int64

var i int64
i = 100
j := i // 最后一行赋值编译会报错; 需要改为强转类型 `j := Duration(i)`
var dur Duration
dur = j
```

## 构造器(一个函数根据数据结构返回这个数据结构的一个实例对象) & 工厂方法

```go
func newStudent(name string, age int) *Student { // 若改为NewStudent(N大写)可以认为是工厂方法
    return &Student{
        name,
        age,
    }
}
```

## 扩展

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
// 可以访问另一个包公开类型的未公开字段(值是类型)的公开字段
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

## 空 struct{}{}

- 不占用内存
- 地址不变


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
