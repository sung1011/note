package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_GetSingletonHungry(t *testing.T) {
	Convey("get instance", t, func() {
		So(GetInstanceHungry(), ShouldEqual, GetInstanceHungry())
	})
}

func Test_GetSingletonLazy(t *testing.T) {
	Convey("get instance", t, func() {
		So(GetInstanceLazy(), ShouldEqual, GetInstanceLazy())
	})
}
