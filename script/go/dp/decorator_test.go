package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Decorator(t *testing.T) {
	Convey("", t, func() {
		Convey("UserDecorator implement order", func() {
			So(UserDecorator{}, ShouldImplement, (*IOrder)(nil))
		})
		Convey("UserDecorator implement user", func() {
			So(UserDecorator{}, ShouldImplement, (*IUser)(nil))
		})
	})
}
