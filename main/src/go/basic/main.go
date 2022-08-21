package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	for i := 0; i < 10; i++ {
		go func(i int) {
			ch <- i
		}(i)
	}
	j := 0
	for {
		if j >= 5 {
			close(ch)
		}
		r, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(r)
		j++
	}
}
