package main

import (
	"bytes"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Slice_Equal(t *testing.T) {
	Convey("", t, func() {
		Convey("equal []byte", func() {
			a := []byte{0, 1, 3, 2}
			b := []byte{0, 1, 3, 2}
			c := []byte{1, 1, 3, 2}
			So(bytes.Equal(a, b), ShouldBeTrue)
			So(bytes.Equal(a, c), ShouldBeFalse)
		})
		Convey("equal []interface{}", func() {
			a := []byte{0, 1, 3, 2}
			b := []byte{0, 1, 3, 2}
			c := []byte{1, 1, 3, 2}
			So(reflect.DeepEqual(a, b), ShouldBeTrue)
			So(reflect.DeepEqual(a, c), ShouldBeFalse)
		})
		Convey("equal func", func() {
			// 自定义比较函数,
			//   优点: 效率更好
			//   缺点: 泛型出来之前, type是被限制的, 每种type都要实现一遍
			f := func(a, b []int) bool {
				// If one is nil, the other must also be nil.
				if (a == nil) != (b == nil) {
					return false
				}
				if len(a) != len(b) {
					return false
				}
				for i := range a {
					if a[i] != b[i] {
						return false
					}
				}
				return true
			}
			a := []int{0, 1, 3, 2}
			b := []int{0, 1, 3, 2}
			c := []int{1, 1, 3, 2}
			So(f(a, b), ShouldBeTrue)
			So(f(a, c), ShouldBeFalse)
		})

	})
}
