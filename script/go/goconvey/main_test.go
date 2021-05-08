package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// install:
//   $ go get github.com/smartystreets/goconvey

// in web-UI:
//   $ goconvey

// in terminal:
//   $ go test
//   $ go test -v

func TestSampleAssert(t *testing.T) { // func name need prefix with Test for testing
	Convey("init data", t, func() { // 顶级Convey调用必须关联*testing.T  top-level calls to Convey(...) need reference to the *testing.T
		x := 1
		Convey("to add", func() {
			x++
			Convey("check val 1", func() {
				So(x, ShouldNotEqual, 1)
			})
			Convey("check val 2", func() { // 同级Convey的title不能相同, 这里不能是check val 1
				So(x, ShouldEqual, 2)
			})
			Convey("check val 3", func() {
				So(x, ShouldNotEqual, 3)
			})
		})
	})
}

func TestSkip(t *testing.T) {
	// SkipConvey 跳过SkipConey和SkipSo的断言
	Convey("skip some", t, func() {
		SkipConvey("skip convey", func() { // skip
			So(1, ShouldEqual, 1)
		})

		Convey("skip so", func() { // skip 1, assert 2
			SkipSo("zxcv", ShouldEqual, "asdf123")
			So("asdf", ShouldEqual, "asdf")
			So("asdf123", ShouldEqual, "asdf123")
		})

		Convey("skip not implemented", nil) // skip

		Convey("normal so", func() { // assert
			So(1, ShouldEqual, 1)
		})
	})
}

func TestFocus(t *testing.T) {
	// FocusConey 会运行嵌套域内所有FocusConvey, 忽略其他断言(是忽略,而非跳过), 调试大型嵌套现场时候用到
	FocusConvey("focus A", t, func() { // assert by focus

		FocusConvey("focus B 1", func() { // assert by focus

			FocusConvey("focus C 1", func() { // assert by focus
				So(1, ShouldEqual, 1)
			})

			Convey("focus C 2", func() { // ignore !
				So(1, ShouldEqual, 1)
			})

		})

		Convey("focus B 2", nil) // ignore !

	})
}

func shouldExactly(actual interface{}, expected ...interface{}) string {
	if actual != "Boo!" {
		return "bad actual"
	}
	if expected[0] != "boo" {
		return "bad expected"
	}
	return ""
}

func TestCustomAssertions(t *testing.T) {
	// 自定义断言
	Convey("custom assertions", t, func() {
		So("Boo!", shouldExactly, "boo")
	})
}
