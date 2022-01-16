package main

import (
	"testing"
	"time"

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
	Convey("", t, func() { // 顶级Convey调用必须关联*testing.T  top-level calls to Convey(...) need reference to the *testing.T
		Convey("general equality", func() {
			So(1, ShouldEqual, 1)
			So(1, ShouldNotEqual, 2)
			So([]byte{97, 98}, ShouldResemble, []byte{97, 98})     // a deep equals for arrays, slices, maps, and structs
			So([]byte{97, 100}, ShouldNotResemble, []byte{97, 98}) // a deep equals for arrays, slices, maps, and structs
			// So(thing1, ShouldPointTo, thing2)
			// So(thing1, ShouldNotPointTo, thing2)
			So(nil, ShouldBeNil)
			So(1, ShouldNotBeNil)
			So(true, ShouldBeTrue)
			So(false, ShouldBeFalse)
			So("", ShouldBeZeroValue)
		})

		Convey("numeric quantity comparison", func() {
			So(1, ShouldBeGreaterThan, 0)
			So(1, ShouldBeGreaterThanOrEqualTo, 0)
			So(1, ShouldBeLessThan, 2)
			So(1, ShouldBeLessThanOrEqualTo, 2)
			So(1.1, ShouldBeBetween, .8, 1.2)
			So(1.1, ShouldNotBeBetween, 2, 3)
			So(1.1, ShouldBeBetweenOrEqual, .9, 1.1)
			So(1.1, ShouldNotBeBetweenOrEqual, 1000, 2000)
			So(1.0, ShouldAlmostEqual, 0.99999999, .0001) // tolerance is optional; default 0.0000000001
			So(1.0, ShouldNotAlmostEqual, 0.9, .0001)
		})

		Convey("collections", func() {
			So([]int{2, 4, 6}, ShouldContain, 4)
			So([]int{2, 4, 6}, ShouldNotContain, 5)
			So(4, ShouldBeIn, []int{2, 4, 6})
			So(4, ShouldNotBeIn, []int{1, 3, 5})
			So([]int{}, ShouldBeEmpty)
			So([]int{1}, ShouldNotBeEmpty)
			So(map[string]string{"a": "b"}, ShouldContainKey, "a")
			So(map[string]string{"a": "b"}, ShouldNotContainKey, "b")
			So(map[string]string{"a": "b"}, ShouldNotBeEmpty)
			So(map[string]string{}, ShouldBeEmpty)
			So(map[string]string{"a": "b"}, ShouldHaveLength, 1) // supports map, slice, chan, and string
		})

		Convey("strings", func() {
			So("asdf", ShouldStartWith, "as")
			So("asdf", ShouldNotStartWith, "df")
			// So("asdf", ShouldEndWith, "df")
			// So("asdf", ShouldNotEndWith, "df")
			So("asdf", ShouldContainSubstring, "sd") // optional 'expected occurences' arguments?
			So("asdf", ShouldNotContainSubstring, "er")
			So("", ShouldBeBlank)
			So("asdf", ShouldNotBeBlank)
		})

		// Convey("panic", func() {
		// 	So(func(), ShouldPanic)
		// 	So(func(), ShouldNotPanic)
		// 	So(func(), ShouldPanicWith, "")    // or errors.New("something")
		// 	So(func(), ShouldNotPanicWith, "") // or errors.New("something")
		// })

		Convey("type checking", func() {
			So(1, ShouldHaveSameTypeAs, 0)
			So(1, ShouldNotHaveSameTypeAs, "asdf")
		})

		Convey("time", func() {
			So(time.Now(), ShouldHappenBefore, time.Now())
			So(time.Now(), ShouldHappenOnOrBefore, time.Now())
			So(time.Now().Add(10*time.Second), ShouldHappenAfter, time.Now())
			So(time.Now().Add(10*time.Second), ShouldHappenOnOrAfter, time.Now())
			So(time.Now(), ShouldHappenBetween, time.Now().Add(-10*time.Second), time.Now().Add(10*time.Second))
			So(time.Now(), ShouldHappenOnOrBetween, time.Now().Add(-10*time.Second), time.Now().Add(10*time.Second))
			So(time.Now(), ShouldNotHappenOnOrBetween, time.Now().Add(-10*time.Second), time.Now().Add(-5*time.Second))
			// So(time.Now(), ShouldHappenWithin, time.Duration.Seconds(), time.Now())
			// So(time.Now(), ShouldNotHappenWithin, time.Duration.Seconds(), time.Now())
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
