package main

import (
	"errors"
	"fmt"
	"os"
	"testing"

	pkgErr "github.com/pkg/errors"

	. "github.com/smartystreets/goconvey/convey"
)

var ErrNotFound = errors.New("not found")

type QueryErr struct {
	Query string
	Err   error
}

func (e *QueryErr) Error() string { return "query err: " + e.Query }

func Test_Errors(t *testing.T) {
	Convey("", t, func() {

		Convey("errors.Is() false", func() {
			So(errors.Is(errors.New("lala"), errors.New("lala")), ShouldBeFalse)
		})
		Convey("errors.Is() true", func() {
			So(errors.Is(fmt.Errorf("prefix: %w", ErrNotFound), ErrNotFound), ShouldBeTrue)
		})

		Convey("errors.As()", func() {
			if _, err := os.Open("non-existing"); err != nil {
				var pathError *os.PathError
				So(errors.As(err, &pathError), ShouldBeTrue)

			}
		})

		Convey("assert", func() {
			_, err := os.Open("non-existing")
			_, ok := err.(*os.PathError)
			So(ok, ShouldBeTrue)
			// So(err, ShouldEqual, os.ErrNotExist)
		})
	})
}

func Test_Errors_lt113(t *testing.T) {
	Convey("errors by assert befor go1.13", t, func() {
		// 实现error接口
		// 存在问题: 无法嵌套, 附加额外信息
		foo := func() error {
			return &QueryErr{"foo", ErrNotFound}
		}
		err := foo()
		if err, ok := err.(*QueryErr); ok && err.Err == ErrNotFound {
			So(fmt.Sprint(err), ShouldEqual, "query err: foo")
		}
	})
}

func Test_Errors_egt113(t *testing.T) {
	Convey("errors by warp after go.1.13 (wrapping error by %w)", t, func() {
		// %w 用来生成一个可包裹的Error的Wrapping Error
		bar := func() error {
			return fmt.Errorf("bar -> %w", ErrNotFound)
		}
		foo := func() error {
			return fmt.Errorf("foo -> %w", bar())
		}
		err := foo()

		So(errors.Is(err, ErrNotFound), ShouldBeTrue)
		So(fmt.Sprint(err), ShouldEqual, "foo -> bar -> not found")
		So(fmt.Sprint(errors.Unwrap(err)), ShouldEqual, "bar -> not found")
		So(fmt.Sprint(errors.Unwrap(errors.Unwrap(err))), ShouldEqual, "not found")

		var err0 *QueryErr
		err1 := fmt.Errorf("new error:[%w]", &QueryErr{"XYZ", errors.New("zzz")})
		So(errors.As(err1, &err0), ShouldBeTrue) // 判断类型是否相同, 并提取第一个符合目标类型的错误.
		So(fmt.Sprint(err0), ShouldEqual, "query err: XYZ")

	})
}

func Test_PkgErrors(t *testing.T) {

	Convey("cause", t, func() {
		e1 := pkgErr.New("inner")
		e2 := pkgErr.Wrap(e1, "middle")
		e3 := pkgErr.Wrap(e2, "outer")

		So(e3.Error(), ShouldEqual, "outer: middle: inner") // with msg, with stack
		So(pkgErr.Cause(e3).Error(), ShouldEqual, "inner")  // 原始error

		e2msg := pkgErr.WithMessage(e2, "with msg")
		So(e2msg.Error(), ShouldEqual, "with msg: middle: inner") // 只with msg, 不with stack
	})

}
