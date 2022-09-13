package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Time(t *testing.T) {
	Convey("", t, func() {
		Convey("time.after", func() {
			before := time.Now()
			c := time.After(1e9)
			after := <-c
			So(after.Unix()-before.Unix(), ShouldEqual, 1)
		})

		var wg sync.WaitGroup
		Convey("timeout", func() {
			ch := make(chan int, 0)
			wg.Add(1)
			go func(ch chan int) {
				defer wg.Done()
				select {
				case v := <-ch:
					t.Log(v)
				case <-time.After(1e9):
					t.Log("timeout") // 先超时
				}
			}(ch)
			go func(ch chan int) {
				time.Sleep(2e9)
				ch <- 123
			}(ch)
			wg.Wait()
			So(true, ShouldBeTrue)
		})
	})
}

func Test_Nil(t *testing.T) {
	Convey("", t, func() {
		Convey("chan = nil", func() {
			ch := make(chan int, 0)
			done := make(chan struct{})
			go func(c chan int) {
				for {
					c <- rand.Intn(10)
				}
			}(ch)

			go func(c chan int, done chan struct{}) {
				sum := 0

				// 1秒后，将向t.C通道发送时间点，使其可读
				tr := time.NewTimer(1 * time.Second)

				for {
					// 一秒内，将一直选择第一个case
					// 一秒后，t.C可读，将选择第二个case
					// c变成nil channel后，两个case分支都将一直阻塞
					select {
					case input := <-c:
						// 不断读取c中的随机数据进行加总
						sum = sum + input
					case <-tr.C:
						c = nil
						t.Log(sum)
						done <- struct{}{}
					}
				}
			}(ch, done)

			<-done
		})
	})
}

func Test_Seq(t *testing.T) {
	Convey("", t, func() {
		Convey("sequence", func() {
			// A首先被a阻塞，A()结束后关闭b，使b可读
			A := func(a, b chan struct{}) {
				<-a
				t.Log("A()!")
				time.Sleep(time.Second)
				close(b)
			}
			// B首先被a阻塞，B()结束后关闭b，使b可读
			B := func(a, b chan struct{}) {
				<-a
				t.Log("B()!")
				close(b)
			}
			// C首先被a阻塞
			C := func(a chan struct{}) {
				<-a
				t.Log("C()!")
			}
			x := make(chan struct{})
			y := make(chan struct{})
			z := make(chan struct{})
			go C(z)
			go A(x, y)
			go C(z)
			go B(y, z)
			go C(z)
			// 关闭x，让x可读
			close(x)
			// 输出顺序 a b c c c
			time.Sleep(3 * time.Second)
		})
	})
}

func Test_gen(t *testing.T) {
	Convey("", t, func() {
		Convey("generate nums", func() {
			generatenums := func(nums chan int) {
				num := 0
				go func() {
					for {
						num++
						nums <- num
					}
				}()
			}
			getnums := func(nums chan int) int {
				return <-nums
			}
			nums := make(chan int)
			generatenums(nums)
			for i := 0; i < 10; i++ {
				t.Log(getnums(nums))
			}

		})
		// 如果多个地方调用generateNums()，各个地方的nums通道将互相影响
		// 所以改为让generateNums()自带属于自己的nums通道，而不是多个generateNums()共享一个nums通道。
		Convey("generate nums fixed", func() {
			generatenums := func() chan int {
				nums := make(chan int)
				num := 0
				go func() {
					for {
						num++
						nums <- num
					}
				}()
				return nums
			}
			getnums := func(nums chan int) int {
				return <-nums
			}
			nums := generatenums()
			for i := 0; i < 10; i++ {
				t.Log(getnums(nums))
			}
		})
	})
}

func Test_Chan_Chan(t *testing.T) {
	Convey("", t, func() {
		Convey("chan chan T", func() {
			// 定义双层通道cc
			cc := make(chan chan int)
			times := 5
			for i := 1; i < times+1; i++ {
				// 定义信号通道f
				f := make(chan bool)
				// 每次循环都在双层通道cc中生成内层通道c
				// 并通过信号通道f来终止f1()
				go func(cc chan chan int, f chan bool) {
					c := make(chan int)
					cc <- c
					defer close(c)
					sum := 0
					select {
					// 从内层通道中取出数据，计算和，然后发回内层通道
					case x := <-c:
						for i := 0; i <= x; i++ {
							sum = sum + i
						}
						// goroutine将阻塞在此，直到数据被读走
						c <- sum
					// 信号通道f可读时，结束f1()的运行
					// 但因为select没有在for中，该case分支用不上
					case <-f:
						return
					}
				}(cc, f)

				// 从双层通道cc中取出内层通道ch
				// 并向ch通道发送数据
				ch := <-cc
				ch <- i

				// 从ch中取出数据
				for sum := range ch {
					t.Logf("Sum(%d)=%d\n", i, sum)
				}
				// 每个循环睡眠一秒钟
				time.Sleep(time.Second)
				// 每次循环都关闭信号通道f
				close(f)
			}
		})
	})

}
