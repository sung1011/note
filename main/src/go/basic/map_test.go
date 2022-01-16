package main

import (
	"fmt"
	"testing"
)

type Foo struct {
	Name string
}

func TestKey(t *testing.T) {
	m := make(map[Foo]int, 10)
	f := Foo{}
	m[f] = 1
	fmt.Println(m)
}
