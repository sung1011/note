package main

import (
	"bytes"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Msg(t *testing.T) {
	Convey("encode", t, func() {
		bf := bytes.NewBuffer([]byte{})

		var err error
		err = encode(bf, "hello world 0")
		So(err, ShouldBeNil)

		err = encode(bf, "hello world 1")
		So(err, ShouldBeNil)

		Convey("decode", func() {
			bs0, err := decode(bf)
			So(err, ShouldBeNil)
			So(string(bs0), ShouldEqual, "hello world 0")

			bs1, err := decode(bf)
			So(err, ShouldBeNil)
			So(string(bs1), ShouldEqual, "hello world 1")
		})
	})
}
