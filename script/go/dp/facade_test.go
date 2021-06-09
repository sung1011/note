package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Facade(t *testing.T) {
	Convey("", t, func() {
		Ifacade := &PC{}
		err, pc := Ifacade.on()
		So(pc, ShouldEqual, &PC{})
		So(err, ShouldBeNil)
	})
}
