# go ptr

## 零值

```go
	var i *int
	fmt.Printf("%p %v\n", i, i) // 0x0 nil
```

## 修改值

```go
func foo(i *int) {
	*i = 100
}

func main() {
	i := new(int)
	foo(i)
	fmt.Println("", *i) // 100
}
```

```go
func swap(a, b *int) {
    // 取a指针的值, 赋给临时变量t
    t := *a
    // 取b指针的值, 赋给a指针指向的变量
    *a = *b
    // 将a指针的值赋给b指针指向的变量
    *b = t

	// *b, *a = *a, *b // 另一种写法
}

func main() {
	x, y := 1, 2
	swap(&x, &y)
	fmt.Println(x, y) // 2 1
}
```