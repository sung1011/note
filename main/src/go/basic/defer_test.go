package main

// 延迟到return后执行
// 倒序执行
// defer func(param), 其参数param会值拷贝一份当前上下文的副本
// 指针和闭包非快照, 会延迟, 取执行时的上下文

// ref: https://www.cnblogs.com/phpper/p/11984161.html

/**
func foo() TYPE{
    ...CODE...
    defer b()
    ...CODE...
    // 函数执行出了错误
    return args
    // 函数b()都会在这里执行
}
**/

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type Users struct {
	name string
}

func (u *Users) GetName() string {
	fmt.Println(u.name)
	return u.name
}
func (u Users) GetNameVal() string {
	fmt.Println(u.name)
	return u.name
}

func GetName(u Users) {
	u.GetName()
}

func Test_Defer_Seq(t *testing.T) {
	// var is []int
	Convey("slice", t, func() {
		var ss [3]struct{} // 0 1 2

		Convey("desc", func() {
			for i := range ss {
				defer fmt.Println(i) // 2 1 0
			}
		})
		Convey("closure 执行时的上下文", func() {
			for i := range ss {
				defer func() {
					fmt.Println(i) // 2 2 2
				}() // 压栈但未传值 执行时的上下文
			}
		})

		Convey("closure(param) 压栈时, 根据上下文传递值拷贝", func() {
			for i := range ss {
				defer func(i int) {
					fmt.Println(i) // 2 1 0
				}(i) // 压栈时, 根据上下文传递值拷贝; 同 `defer fmt.Println(i)`
			}
		})
		Convey("return; 闭包延迟读取, 带命名返回值", func() {
			// 模拟普通函数 带有命名返回值
			func() (s string) {
				s = "aaa"
				defer func() {
					fmt.Println("defer: " + s) // derfer: bbb
				}()
				return "bbb"
				// defer s := bbb
			}()
		})

	})
}

func Test_Defer_Panic(t *testing.T) {
	Convey("panic", t, func() {
		Convey("seq", func() { // b a panic; (panic 后的 c 不会被执行到), panic最后执行
			defer t.Log("a")
			defer t.Log("b")
			// panic("errrrr") // 当前与return之后
			defer t.Log("c") // be ignore
		})
		Convey("closure", func() { // d b a panic; (不会执行到c), panic最后执行
			defer t.Log("a")
			defer func() {
				t.Log("b")
				// panic("errrrrr")
				t.Log("c") // be ignore
			}()
			defer t.Log("d")
		})
		Convey("recover", func() { // b recover a; (panic 后的 c 不会被执行到)
			defer t.Log("a")
			defer func() {
				t.Log("b")
				if err := recover(); err != nil {
					t.Log("recover")
				}
			}()
			panic("errrrr")
			defer t.Log("c")
		})
	})
}

func Test_Defer_Struct(t *testing.T) {
	Convey("struct", t, func() {
		Convey("指针接收者 => last", func() {
			list := []Users{{"a"}, {"b"}, {"c"}}
			for _, u := range list {
				defer u.GetName() // c c c
			}
		})

		Convey("指针接收者(param) => ok", func() {
			list := []Users{{"a"}, {"b"}, {"c"}}
			for _, u := range list {
				defer GetName(u) // c b a
			}
		})

		Convey("值接收者 => ok", func() { // copy
			list := []Users{{"a"}, {"b"}, {"c"}}
			for _, u := range list {
				defer u.GetNameVal() // c b a
			}
		})
	})
}

func Test_Defer_Return(t *testing.T) {
	Convey("return", t, func() {
		Convey("尽量不在循环中使用defer, defer有额外开销; 去掉defer即可", func() {
			for i := 0; i < 100; i++ {
				// f, _ := os.Open("./defer_test")
				// defer f.Close()
			}
		})

		Convey("判断err之后, 再defer释放资源; 有err则没必要defer了, 直接返回", func() {
			// f, err := os.Open("./defer_test")
			// if err != nil {
			// 	return err
			// }
			// defer f.Close()
		})

		Convey("os.Exit时defer不会执行", func() {
			defer func() {
				fmt.Println("defer...")
			}()
			// os.Exit(0)
		})
	})
}

/**
Benchmark_WithDefer-4      	 3196197	       358.1 ns/op
Benchmark_WithoutDefer-4   	1000000000	         0.4087 ns/op
*/

func Benchmark_WithDefer(b *testing.B) {
	var j int = 0
	for i := 0; i < b.N; i++ {
		defer func() {
			j++
		}()
	}
}

func Benchmark_WithoutDefer(b *testing.B) {
	var j int = 0
	for i := 0; i < b.N; i++ {
		j++
	}
}
