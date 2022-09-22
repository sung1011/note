package main

import (
	"fmt"
	"sync"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func GoroutineNotPanic(tasks ...func() error) (err error) {
	var wg sync.WaitGroup
	for _, f := range tasks { // 模拟并发
		wg.Add(1)
		go func(f func() error) {
			defer func() {
				// 每个协程内部使用recover捕获可能在调用逻辑中发生的panic
				if e := recover(); e != nil {
					fmt.Println("panic", e)
				}
				wg.Done()
			}()
			// 取第一个报错的task调用逻辑，并最终向外返回
			e := f()
			if err == nil && e != nil {
				err = e
			}
		}(f)
	}
	wg.Wait()
	return
}

func Test_Panic(t *testing.T) {
	Convey("", t, func() {
		Convey("recover可以取消panicking状态", func() {
			t.Log("11111")
			func() {
				defer func() {
					if e := recover(); e != nil {
						t.Log("panic is:", e)
					}
				}()
				panic("xxxxx")
				t.Log("no output") // panicking状态, 无法执行
			}()
			t.Log("22222") // 由于func中的recover()已经取消了panicking, 故可以执行到22222
		})

		Convey("单测shouldPanic\n", func() {
			f1 := func() {
				panic("paniccc")
			}
			So(func() { f1() }, ShouldPanic)
		})

		Convey("g1 panic, g2跟着挂掉\n", func() {
			f1 := func() {
				go func() {
					// panic("paniccc") // 会panic Convey原因接不住所以屏蔽
				}()
				time.Sleep(1e9)
			}
			f1()
			// So(func() { f1() }, ShouldPanic)
		})

		Convey("g1 panic, g2的recover接不到\n", func() {
			f2 := func() {
				defer func() {
					if e := recover(); e != nil {
						fmt.Println("got!") // 接不到
					}
				}()
				go func() {
					// panic("paniccc")
				}()
				time.Sleep(1e9)
			}
			f2()
			// So(func() { f2() }, ShouldPanic)
		})
		Convey("每个goroutine都要有recover()\n", func() { // gin有类似实现
			aRpc := func() error {
				panic("rpc logic A panic")
				return nil
			}

			bRpc := func() error {
				fmt.Println("rpc logic B")
				return nil
			}

			err := GoroutineNotPanic(aRpc, bRpc)
			if err != nil {
				fmt.Println(err)
			}
		})

		Convey("触发多个panic, recover最近的一个\n", func() {
			defer func() {
				t.Log("defer_0") // 会输出
			}()
			// defer func() {
			// 	panic("panic 3") // 不会被recover, 因为实在recover()之后执行
			// }()
			defer func() {
				if e := recover(); e != nil {
					t.Log("recovered", e) // 输出最近一个 即panic 2
				}
			}()
			defer func() {
				panic("panic 2")
			}()
			panic("panic 1")
		})

	})
}
