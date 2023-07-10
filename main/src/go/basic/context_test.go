package main

import (
	"context"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func watchContext(ctx context.Context, ch chan<- int32, name int32) {
	for {
		select {
		case <-ctx.Done():
			ch <- name
			return
		}
	}
}

func Test_Context(t *testing.T) {
	Convey("", t, func() {
		ctx, cancel := context.WithCancel(context.Background())
		ch := make(chan int32)
		go watchContext(ctx, ch, 111)
		go watchContext(ctx, ch, 222)
		go watchContext(ctx, ch, 333)

		time.Sleep(2 * time.Second)
		cancel()
		bs := make([]int32, 0, 3)
		for {
			v := <-ch
			So(v, ShouldBeIn, []int32{111, 222, 333})
			bs = append(bs, v)
			if len(bs) >= 3 {
				break
			}
		}
	})
}
