# 1. go 结构体

- [1. go 结构体](#1-go-结构体)
  - [1.1. 数据结构](#11-数据结构)
  - [1.2. 创建 初始化 访问](#12-创建-初始化-访问)
  - [1.3. 构造函数](#13-构造函数)
  - [1.4. 内存布局](#14-内存布局)
  - [1.5. 参数传递](#15-参数传递)
  - [1.6. 是否==](#16-是否)
  - [1.7. 方法](#17-方法)
  - [1.8. 包含 \& 组合 \& 重载 \& 内嵌](#18-包含--组合--重载--内嵌)
  - [1.9. 标签tag](#19-标签tag)
  - [1.10. 嵌套接口](#110-嵌套接口)
  - [1.11. 未公开类型的公开字段](#111-未公开类型的公开字段)
  - [1.12. 空结构体 struct{}{}](#112-空结构体-struct)
  - [1.13. clone](#113-clone)
  - [1.14. 导入导出](#114-导入导出)
  - [1.15. 常用的struct](#115-常用的struct)
  - [1.16. ref](#116-ref)

## 1.1. 数据结构

        TODO

## 1.2. 创建 初始化 访问

```go
type Student struct {// 结构体
    Name string
    Age int
    Scores map[string]int
}
// 声明; 已初分配内存和初始化为零值; 零值可用类型. 如: var b bytes.Buffer, var mu sync.Mutex
var s Student 
s.Name = "Sam"
s.Age = 18

// 初始化
sp1 := new(Student) // 等价于 &Student{} 

// 初始化 并赋值
sp2 := &Student{Name: "Sam", Age: 18, Scores: make(map[string]int)}
sp2.Scores = map[string]int{"math": 98, "chemistry": 80}

// 初始化 并顺序赋初始值; 不推荐
sp3 := &Student{"Sam", 18} 
```

## 1.3. 构造函数

    作为工厂方法用于构造实例, 返回实例的指针
    特别是可以用来对未导出的字段赋值
    与new()效果完全相同

```go
// func NewT(field1, field2, ...) *T
func NewStudent(name string, age int) *Student {
    return &Student{
        name: name,
        age: age,
    }
}
```

## 1.4. 内存布局

    `连续` 结构体和它所包含的数据在内存中是以连续块的形式存在的, 即使结构体中嵌套有其他的结构体
    `对齐` 每个字段至少占用8byte, 字段顺序影响其内存空间占用

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
	fmt.Println(unsafe.Sizeof(t1), unsafe.Sizeof(t2)) // 24 16
}
```

![img](res/go-struct-assign.jpg)

![img](res/go-struct-mem.jpg)

## 1.5. 参数传递

    一般传递其指针 &struct (ptr 8byte)
    不传指针, 则是结构体本身, 参考内存布局

> [内存布局](go-struct.md#内存布局)

## 1.6. [是否==](go-type-compare.md#struct)

## 1.7. 方法

    属于结构体的函数

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

> recv是值类型or引用类型, 引用类型用值接收者也可修改实例本身(因为是传递指针的副本), 值类型区分区分是修改实例本身or生成副本

> 如果一个工厂函数返回的是指针, 其方法都应当使用指针接收者 (即使该方法不修改实例本身); 例外: 需要类型值符合某个接口时

> func(recv) 是函数; recv.Method() 是方法

## 1.8. [包含 & 组合 & 重载 & 内嵌](go-struct-relation.md)

## 1.9. 标签tag

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

## 1.10. [嵌套接口](go-interface.md#嵌入类型)

## 1.11. 未公开类型的公开字段

```go
    TODO
```

## 1.12. 空结构体 struct{}{}

    不占用内存
    地址不变

```go
    type Empty struct{}
    var s Empty
    fmt.Println(unsafe.Sizeof(s)) // 0

    var ch := make(chan Empty)
    ch <-Empty  // 不占用内存, 只作为一种事件信息传递
```

## 1.13. clone

    todo

## 1.14. [导入导出](go-import.md#struct)

## 1.15. [常用的struct](go-struct-frequently.md)

## 1.16. ref

- <https://www.cnblogs.com/myuniverse/p/11595043.html>
