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

sp := &Student{Name:"Sam", "Age":18} // 初始化 与new等效
```

![img](res/go-struct-assign.jpg)

## 内存布局 连续

结构体和它所包含的数据在内存中是以连续块的形式存在的，即使结构体中嵌套有其他的结构体

![img](res/go-struct-mem.jpg)

## 参数传递

### 问题

1. 函数内部无法修改传递给函数的原始struct，它修改的是拷贝后的副本。
2. 副本可能很大。

### 解决

一般struct作为参数时，传递的都是引用/指针( &struct )

## 构造器（一个函数根据数据结构返回这个数据结构的一个实例对象） & 工厂方法

```go
func newStudent(name string, age int) *Student { // 若改为NewStudent（N大写）可以认为是工厂方法
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

bm_horse := &Animal{
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
    name   string
    weight int
}

type Horse struct {
    *Animal //组合, 内嵌
    speak string
    int //匿名字段
}

func (a *Animal) hello() {
    fmt.Println(a.name)
    fmt.Println(a.weight)
    //fmt.Println(a.speak) // animal访问不到horse的字段， 即便是horse实例
    // fmt.Println(a.int) // animal访问不到horse的匿名字段
}

func (h *Horse) hello() { // 重载 覆盖 Animal.hello()
    fmt.Println(h.name)
    fmt.Println(h.weight)
    fmt.Println(h.speak)
    fmt.Println(h.int)
}

func main() {
    bm_horse := &Horse{
        Animal: &Animal{ // 结构体包含**嵌入字段**，类型名充当了嵌入字段的字段名
            name:   "baima",
            weight: 60,
        },
        speak: "neigh",
        int: 1234,
    }
    bm_horse.hello() // horse.hello()
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

## 函数func(recv) & 方法recv.Method()

## [嵌套接口](go-interface.md#嵌入结构体)
