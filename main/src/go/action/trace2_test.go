package main

import (
	"fmt"
	"runtime"
	"testing"
	// . "github.com/smartystreets/goconvey/convey"
)

/**
trace函数 不用显示的传递参数了
*/

func trace2() func() {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("not found caller")
	}
	name := runtime.FuncForPC(pc).Name()
	fmt.Println("[in]", name)
	return func() {
		fmt.Println("[out]", name)
	}
}

func foo2() {
	defer trace2()()
	bar2()
}

func bar2() {
	defer trace2()()
}

func Test_Trace2(t *testing.T) {
	foo2()
}
