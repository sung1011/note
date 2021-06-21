package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Template_Method(t *testing.T) {
	Convey("", t, func() {
		xhs := &XiHongShi{}
		doCook(xhs)
		So(xhs.Dish, ShouldEqual, "我是一盘西红柿")

		cjd := &ChaoJiDan{}
		doCook(cjd)
		So(cjd.Dish, ShouldEqual, "我是一盘炒鸡蛋")
	})
}
