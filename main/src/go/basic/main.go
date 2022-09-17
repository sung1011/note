package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 0)
	go func() {
		for {
			ch <- 1
		}
	}()

	t := time.NewTimer(1e9)
	for {
		select {
		case v := <-ch:
			fmt.Println("", v)
		case <-t.C:
			fmt.Println("timeout")
			return
		}
	}
}
