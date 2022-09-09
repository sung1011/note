# go type 是否可比较

    引用类型可使用==nil比较
    值类型不可使用==nil比较

## array

    len相同 可==; len不同 不可== 

> 复合类型

> [3]int, [4]int 是两种类型, 不可==

## struct

    由于struct是复合类型
    只包含可比较T 可==; 包含不可比较T 不可==

> 复合类型

## int

## float

> 注意精度

## string

## bool

## slice

    不可==

> [判断相等](src/go/basic/slice_test.go)

## map

    不可==

> map的key 是可比较T, 则可比较;

## func

    不可==

## chan

    判断内存地址

```go
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := ch1
	fmt.Println(ch1 == ch2) // false
	fmt.Println(ch1 == ch3) // true
```

## ptr

    判断内存地址

```go
	i1, i2 := 1, 1
	var i3, i4 *int
	fmt.Println(&i1 == &i2, i3 == i4) // false, true
```

## interface

    类型, 值 都相等, 接口才相等

```go
	var a interface{} = 0
	var b interface{} = 2
	var c interface{} = 0
	var d interface{} = 0.0
	fmt.Println(a == b) // false
	fmt.Println(a == c) // true
	fmt.Println(a == d) // false
```

```go
type Speaker interface {
	Speak()
}

type Person struct {
	name string
}

func (p Person) Speak() {
	fmt.Println(p.name)
}

type Student struct {
	name string
}

func (s Student) Speak() {
	fmt.Println(s.name)
}

func main() {
	fmt.Println(
        // 接口==接口
		Speaker(Person{"ball"}) == Speaker(Person{"ball"}),  // true
		Speaker(Person{"ball"}) == Speaker(Person{"x"}),     // false
		Speaker(Person{"ball"}) == Speaker(Student{"ball"}), // false
		// 接口==非接口
		Speaker(Person{"ball"}) == Person{"ball"},  // true
		Speaker(Person{"ball"}) == Person{"x"},     // false
		Speaker(Person{"ball"}) == Student{"ball"}, // false
		// 接口==nil
		Speaker(Student{"ball"}) == nil, // false
	)
}
```

## ref

- 类型比较 <https://segmentfault.com/a/1190000039005467>
- <https://cloud.tencent.com/developer/article/1792403>