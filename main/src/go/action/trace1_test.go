package main

import (
	"fmt"
	"testing"
	// . "github.com/smartystreets/goconvey/convey"
)

func trace1(name string) func() {
	fmt.Println("[in]", name)
	return func() {
		fmt.Println("[out]", name)
	}
}

func foo1() {
	defer trace1("foo")()
	bar1()
}

func bar1() {
	defer trace1("bar")()
}

func Test_Trace1(t *testing.T) {
	foo1()
}
