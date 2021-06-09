package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Composite(t *testing.T) {
	Convey("", t, func() {
		Convey("the Employee num", func() {
			So(NewOrganization().Count(), ShouldEqual, 12)
		})
	})
}
