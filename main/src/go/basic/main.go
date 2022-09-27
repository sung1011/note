package main

import (
	"fmt"
)

func main() {
	fmt.Println("foo1")
	bar()
	fmt.Println("foo2")
}

func bar() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("panic is:", e)
		}
	}()
	fmt.Println("bar1")
	panic("barrrr")
	fmt.Println("bar2")
}
