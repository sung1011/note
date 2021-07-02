package main

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_File(t *testing.T) {
	Convey("", t, func() {

		o := New(tmp_file)

		Convey("create empty file", func() {
			Convey("file not exist", func() {
				So(o.Exist(), ShouldBeFalse)
				Convey("create()", func() {
					_, err := o.Create()
					So(err, ShouldBeNil)
					Convey("file exist", func() {
						So(o.Exist(), ShouldBeTrue)
					})
				})
			})
			defer o.Remove()
		})

		Convey("truncate", func() {
			content := "abcdefg"
			Convey("init file -> "+content, func() {
				err := o.WriteSimple([]byte(content))
				So(err, ShouldBeNil)
				Convey(fmt.Sprintf("truncate(3) %s -> abc", content), func() {
					err = o.Truncate(3)
					So(err, ShouldBeNil)
					bs, err := o.ReadSimple()
					So(err, ShouldBeNil)
					So(bs, ShouldResemble, []byte("abc"))

				})
				Convey(fmt.Sprintf("truncate(10) %s -> %s<nil><nil><nil>", content, content), func() {
					err = o.Truncate(10)
					So(err, ShouldBeNil)
					bs, err := o.ReadSimple()
					So(err, ShouldBeNil)
					// t.Log(bs, len(bs), string(bs))
					rs := append([]byte(content), 0, 0, 0)
					So(bs, ShouldResemble, rs)
				})

				Convey(fmt.Sprintf("truncate(0) %s -> ", content), func() {
					err = o.Truncate(0)
					So(err, ShouldBeNil)
					bs, err := o.ReadSimple()
					So(err, ShouldBeNil)
					So(bs, ShouldBeEmpty)
				})
			})

			defer o.Remove()
		})
	})
}
