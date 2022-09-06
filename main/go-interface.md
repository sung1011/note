# go 接口

        相同行为, 不同实例 (多态)

## 数据结构

        TODO

## 方法集methodSet & 参数

### 概述

        接口指针类型接收者`recv(*T)`只能接收指针实例`*T`, 不能接收值类型实例`T` (注意区别接口方法 和 普通方法, 普通方法无此限制)

### 原理

        实例的methodSet决定了它所实现的接口, 以及通过receiver可以调用的方法.  
        通过指针实例可以拿到值类型实例的methodSet(解引用), 通过值实例不能拿到指针实例的methodSet

> 多态

### 示例

```go
type I interface {
	M1()
	M2()
}

type Foo struct {}

func (*Foo) M1() {      // 指针接收者
	fmt.Println("m1")
}
func (Foo) M2() {       // 值接收者
	fmt.Println("m2")
}

func main() {
	var i I
	i = &Foo{}
	// i = Foo{}   // error; 值 没实现接口I的M1方法
	i.M1() // m1
	i.M2() // m2
    _, ok1 := i.(I)      // 接口断言, 实例是否实现了I接口
	_, ok2 := i.(*Foo)
	// _, ok2 := i.(Foo)    // error; 值 没实现接口I的M1方法
	fmt.Println("implement:", ok1, ok2) // true true
}

// 接口 用于同一行为, 实例替换, 如 db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)") // mysql可换成sqllite, postgresql, oralce  

// 接口 也用于不同类型统一行为, 如 copy(os.Stdout, r.Body)  
```

### 接收者角度

| receiver  | 实例的类型 | 注解                                                    |
| --------- | ---------- | ------------------------------------------------------- |
| (T Type)  | T 或 *T    | 实现接口的方法的接收者是值类型更通用                   |
| (T \*Type) | *T         | `实现接口的方法的接收者是指针类型, 只能接收指针实例` |

### 实例角度

| 实例的类型  | receiver            | 注解                                 |
| ----------- | ------------------- | ------------------------------------ |
| 值类型:T    | (T Type)            | 接口中实例是值类型, 只能组成值接收者 |
| 指针类型:\*T | (T Type)或(T *Type) | 接口中实例是指针类型更通用           |

## 嵌套 组合

### 接口组合

```go
type IReader interface {
   Read(file string) []byte
}

type IWriter interface {
   Write(file string, data string)
}

// 默认继承了IReader和IWriter中的抽象方法
// 方法需要实现IReader 和 IWriter中全部接口方法.
type IReadWriter interface {
   IReader
   IWriter
}
```

### 嵌入类型

```go
type I interface {
    do()
}

type Foo struct {
    Name string
    I    // 匿名接口字段; 嵌入结构体
}

type Inner struct {
    int
}

func (in Inner) do() {}

func main() {
    f := Foo{Name: "abc", I: &Inner{3}} // 值类型实例Inner{3}也可以
}
```

### 方法提升

1. 如果S包含一个匿名字段T, S和S的方法集都包含接收者为T的方法提升.

   当嵌入一个类型, 嵌入类型的接收者为值类型的方法将被提升, 可以被外部类型的值和指针调用.

2. 对于S类型的方法集包含接收者为T的方法提升

   当嵌入一个类型, 可以被外部类型的指针调用的方法集只有嵌入类型的接收者为指针类型的方法集, 即当外部类型使用指针调用内部类型的方法时, 只有接收者为指针类型的内部类型方法集将被提升.

3. 如果S包含一个匿名字段T, S和S的方法集都包含接收者为T或者T的方法提升

   当嵌入一个类型的指针, 嵌入类型的接收者为值类型或指针类型的方法将被提升, 可以被外部类型的值或者指针调用.

4. 如果S包含一个匿名字段T, S的方法集不包含接收者为*T的方法提升.

   根据Go语言规范里方法提升中的三条规则推导出的规则.当嵌入一个类型, 嵌入类型的接收者为指针的方法将不能被外部类型的值访问.

## 空接口 interface{}

      可存储任意数据类型的实例

```go
type Foo struct {
    Is []interface{} // 可混合存储任意类型元素
}

func main() {
    var f Foo
    is := make([]interface{}, 4)
    f.Is = is
    f.Is[0] = 123
    f.Is[1] = "a"
    f.Is[2] = []int{3}
    fmt.Println(f) // {[123 a [3] nil]}
}
```

## 类型断言

```go
var x interface{}
x = 10
v, ok := x.(int) // 10, true; 断言x是否为实现了int类型(的实例10)的接口.
// v, ok := x.(*int) // nil, false; 断言有严格的判断.

// 配合switch
switch x.(type) {
    case int:
       fmt.Println("is int")
    case string:
       fmt.Println("is string")
    default:
       fmt.Println("unknown type")
}
```

> 接口类型才可以断言, 如果x:=10, 需要转换为接口才能断言: switch interface{}(x).(type)

## ref

- [Go语言开发(五)、Go语言面向接口](https://blog.51cto.com/9291927/2130244)
