package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Singleton_Hungry(t *testing.T) {
	Convey("get instance", t, func() {
		So(GetInstanceHungry(), ShouldEqual, GetInstanceHungry())
	})
}

func Test_Singleton_Lazy(t *testing.T) {
	Convey("get instance", t, func() {
		So(GetInstanceLazy(), ShouldEqual, GetInstanceLazy())
	})
}
