package main

import (
	"fmt"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Create(t *testing.T) {
	Convey("create empty file", t, func() {

		o := New(tmp_file)

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
}

func Test_Truncate(t *testing.T) {
	Convey("truncate file content", t, func() {

		o := New(tmp_file)

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
}

func Test_Permission(t *testing.T) {
	Convey("check permission", t, func() {

		o := New(tmp_file)

		Convey("create file", func() {
			_, err := o.Open(os.O_CREATE|os.O_RDONLY|os.O_TRUNC, 0666)
			So(err, ShouldBeNil)

			Convey("chmod rw -> r", func() {
				err := o.Chmod(0444)
				So(err, ShouldBeNil)
				Convey("check permission write", func() {
					_, err := o.Open(os.O_WRONLY, 0666)
					So(os.IsPermission(err), ShouldBeTrue)

					Convey("chown", func() {
						err := o.Chown(os.Getuid(), os.Getgid())
						// TODO 如何获取文件的own, group?
						So(err, ShouldBeNil)
					})
				})
			})
		})
		defer o.Remove()
	})
}

func Test_SymLink(t *testing.T) {
	Convey("", t, func() {

	})
}
