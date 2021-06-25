package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Mediator(t *testing.T) {
	Convey("", t, func() {
		Convey("login", func() {
			selection := Selection("登录")
			un := Input("username input")
			pwd := Input("password input")
			rpwd := Input("repeatPwd input")
			d := &Dialog{
				Selection:           &selection,
				UsernameInput:       &un,
				PasswordInput:       &pwd,
				RepeatPasswordInput: &rpwd,
			}
			So(d.HandleEvent(&selection), ShouldEqual, 3)

			Convey("reg", func() {
				regSelection := Selection("注册")
				d.Selection = &regSelection
				So(d.HandleEvent(&regSelection), ShouldEqual, 2)
			})
		})
	})
}
