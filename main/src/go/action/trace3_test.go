package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"

	// . "github.com/smartystreets/goconvey/convey"
	"github.com/petermattis/goid"
)

/**
trace函数 不用显示的传递参数了
并发时, 能看到goid
*/
func trace3() func() {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("not found caller")
	}
	name := runtime.FuncForPC(pc).Name()
	goid := goid.Get()
	fmt.Printf("[-> g%05d] %s\n", goid, name)
	return func() {
		fmt.Printf("[<- g%05d] %s\n", goid, name)
	}
}

func foo3() {
	defer trace3()()
	bar3()
	time.Sleep(1e9)
}

func bar3() {
	defer trace3()()
}

func Test_Trace3(t *testing.T) {
	var wg sync.WaitGroup
	var count = 2
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			foo3()
			wg.Done()
		}()
	}
	wg.Wait()
}
