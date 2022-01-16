package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Strategy(t *testing.T) {
	Convey("", t, func() {
		Convey("choice strategy", func() {
			enc, err := NewStorageStrategy("encrypt_file")
			So(err, ShouldBeNil)
			Convey("exec strategy", func() {
				err = enc.Save("foo", []byte("abc"))
				So(err, ShouldBeNil)
			})
		})
	})
}
