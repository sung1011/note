package main

// 延迟到return后执行
// 倒序执行
// defer func(param), 其参数param会快照保存一份(defer上下文的)副本
// 指针和闭包会延迟

// ref: https://www.cnblogs.com/phpper/p/11984161.html

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

func Test_Defer(t *testing.T) {
	// var is []int
	Convey("", t, func() {
		Convey("slice", func() {
			var ss [3]struct{} // 0 1 2

			Convey("_; desc", func() { // 2 1 0
				for i := range ss {
					defer fmt.Println(i)
				}
			})
			// 闭包用到的变量, 执行时已经变成了2
			Convey("closure; last; ", func() { // 2 2 2
				for i := range ss {
					defer func() {
						fmt.Println(i)
					}()
				}
			})

			Convey("func ; desc", func() { // 2 1 0
				for i := range ss {
					defer func(i int) {
						fmt.Println(i)
					}(i)
				}
			})

		})

		Convey("struct", func() {
			Convey("指针接收者; last", func() { // c c c
				list := []Users{{"a"}, {"b"}, {"c"}}
				for _, u := range list {
					defer u.GetName()
				}
			})

			Convey("by 参数; desc", func() { // c b a
				list := []Users{{"a"}, {"b"}, {"c"}}
				for _, u := range list {
					defer GetName(u)
				}
			})

			Convey("值接收者; desc", func() { // c b a
				list := []Users{{"a"}, {"b"}, {"c"}}
				for _, u := range list {
					defer u.GetNameVal()
				}
			})
		})

		Convey("panic", func() {
			Convey("defer内容执行后才会报错", func() { // d c b a panic
				defer fmt.Println("a")
				defer fmt.Println("b")

				i := 0
				defer func(i int) {
					fmt.Println("c")
					// fmt.Println(10 / i) // panic
					// panic("errrrrr") // panic
				}(i)

				defer fmt.Println("d")
			})
		})

		Convey("return; 注意命名返回值; 闭包延迟读取", func() { // derfer: b
			// 模拟普通函数 带有命名返回值
			func() (s string) {
				s = "aaa"
				defer func() {
					fmt.Println("defer: " + s)
				}()
				return "bbb" // 会赋值给命名返回值s
			}()
		})

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
