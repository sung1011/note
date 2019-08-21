# go 接口

## 数据结构

## 方法集methodSet & 参数

### 概述

接口指针类型接收者`recv(*T)`只能接收指针实例`*T`，不能接收值类型实例`T`。(注意区别接口方法 和 普通方法，普通方法无此限制)

### 原理

实例的method set决定了它所实现的接口，以及通过receiver可以调用的方法。  
通过指针实例可以拿到值类型实例的methodSet(解引用)，通过值实例不能拿到指针实例的methodSet。  

### 示例

```go
type I interface {
    do()
}

func (t *T) do() { } // 指针接收者

func main() {
    var p, v I

    p = &T{} // 实现接口的指针实例
    p.do() // 指针实例ok

    v = T{} // 实现接口的值实例
    // v.do() // 值实例error; 值实例找不到其接口指针*T的方法集

    foo(T{}) // 作为参数 传递实例
}

func foo(i I) { // 作为参数 转为接口
    i.do()
}
```

### 接收者角度

| receiver  | 实例的类型 | 注解                                                    |
| --------- | ---------- | ------------------------------------------------------- |
| (T Type)  | T 或 *T    | 实现接口的方法 的接收者是值类型更通用                   |
| (T *Type) | *T         | **实现接口的方法 的接收者是指针类型，只能接收指针实例** |

### 实例角度

| 实例的类型  | receiver            | 注解                                 |
| ----------- | ------------------- | ------------------------------------ |
| 值类型:T    | (T Type)            | 接口中实例是值类型，只能组成值接收者 |
| 指针类型:*T | (T Type)或(T *Type) | 接口中实例是指针类型更通用           |

## 嵌套 组合

### 组合

接口组合: 默认继承了IReader和IWriter中的抽象方法；方法需要实现IReader, IWriter中全部接口方法。

```go
type IReader interface {
   Read(file string) []byte
}

type IWriter interface {
   Write(file string, data string)
}

type IReadWriter interface {
   IReader
   IWriter
}
```

### 嵌入结构体

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

1. 如果S包含一个匿名字段T，S和S的方法集都包含接收者为T的方法提升。
当嵌入一个类型，嵌入类型的接收者为值类型的方法将被提升，可以被外部类型的值和指针调用。
2. 对于S类型的方法集包含接收者为T的方法提升
当嵌入一个类型，可以被外部类型的指针调用的方法集只有嵌入类型的接收者为指针类型的方法集，即当外部类型使用指针调用内部类型的方法时，只有接收者为指针类型的内部类型方法集将被提升。
3. 如果S包含一个匿名字段T，S和S的方法集都包含接收者为T或者T 的方法提升
当嵌入一个类型的指针，嵌入类型的接收者为值类型或指针类型的方法将被提升，可以被外部类型的值或者指针调用。
4. 如果S包含一个匿名字段T，S的方法集不包含接收者为*T的方法提升。
根据Go语言规范里方法提升中的三条规则推导出的规则。当嵌入一个类型，嵌入类型的接收者为指针的方法将不能被外部类型的值访问。

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

语法

```go
var x interface{}
x = 10
v, ok := x.(int) // 断言x是否为实现了int类型(的实例10)的接口。 10, true
// v, ok := x.(*int) // 断言有严格的判断。 nil, false
```

配合switch

```go
switch x.(type) {
    case int:
    fmt.Println("the type of a is int")
    case string:
    fmt.Println("the type of a is string")
    default:
    fmt.Println("unknown type")
}
```

[tag](go-struct.md#标签tag)

## ref

[Go语言开发（五）、Go语言面向接口](https://blog.51cto.com/9291927/2130244)
