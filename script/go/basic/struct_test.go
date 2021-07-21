package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type Struct_Foo struct {
	Name string
}

func Test_Struct_Eq(t *testing.T) {
	foo := Struct_Foo{}
	bar := Struct_Foo{}
	Convey("equal", t, func() {
		So(foo == bar, ShouldBeTrue)

		So(&foo != &bar, ShouldBeTrue)
	})
}
