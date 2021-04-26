package main

import (
	"testing"
)

type Struct_Foo struct {
	Name string
}

func TestStruct_Eq(t *testing.T) {
	foo := Struct_Foo{}
	bar := Struct_Foo{}
	if foo != bar {
		t.Error("not eq")
	}
}
