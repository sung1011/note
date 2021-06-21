package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Iterator(t *testing.T) {
	Convey("for each", t, func() {
		d := ArrayInt{1, 3, 5, 7, 8}
		iterator := d.Iterator()
		i := 0
		for iterator.HasNext() {
			So(iterator.CurrentItem(), ShouldEqual, d[i])
			iterator.Next()
			i++
		}
	})
}
