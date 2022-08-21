package main

// 延迟到return后执行
// 倒序执行
// defer func(param), 其参数param会快照保存一份(defer上下文的)副本
// 指针和闭包非快照, 会延迟, 取执行时的上下文

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

func Test_Defer_Slice(t *testing.T) {
	// var is []int
	Convey("slice", t, func() {
		var ss [3]struct{} // 0 1 2

		Convey("=> ok", func() {
			for i := range ss {
				defer fmt.Println(i) // 2 1 0
			}
		})
		// 闭包用到的变量, 执行时已经变成了2
		Convey("closure => last; ", func() {
			for i := range ss {
				defer func() {
					fmt.Println(i) // 2 2 2
				}()
			}
		})

		Convey("closure(param) => ok", func() {
			for i := range ss {
				defer func(i int) {
					fmt.Println(i) // 2 1 0
				}(i)
			}
		})

	})
}

func Test_Defer_Panic(t *testing.T) {
	Convey("panic", t, func() {
		Convey("seq", func() { // b a panic (panic 后的 c 不会被执行到)
			defer t.Log("a")
			defer t.Log("b")
			// panic("errrrr")
			defer t.Log("c")
		})
		Convey("closure", func() { // c b a panic (会执行到c)
			defer t.Log("a")
			pf := func() {
				t.Log("b")
				// panic("errrrrr") // panic
			}
			defer pf()
			defer t.Log("c")
		})
		Convey("recover", func() { // b recover a (panic 后的 c 不会被执行到)
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

func Test_Defer_Over(t *testing.T) {
	Convey("over", t, func() {
		Convey("panic => defer内容执行后才会报错", func() { // d c b a panic
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

		Convey("return; 注意命名返回值; 闭包延迟读取", func() {
			// 模拟普通函数 带有命名返回值
			func() (s string) {
				s = "aaa"
				defer func() {
					fmt.Println("defer: " + s) // derfer: bbb
				}()
				return "bbb" // 会先返回 赋值给命名返回值s
			}()
			// 相当于
			// 1. s=aaa
			// 2. return bbb
			// 3. func() (s=bbb) {}
			// 4. defer s=bbb
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
