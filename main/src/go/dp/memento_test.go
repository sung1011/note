package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Memento(t *testing.T) {
	Convey("", t, func() {
		in := &InputText{}
		in.Append("aaa")
		So(in.GetText(), ShouldEqual, "aaa")

		Convey("snap independ on input", func() {
			snap := in.Snapshot()
			in.Append("bbb")
			So(in.GetText(), ShouldEqual, "aaabbb")
			So(snap.GetText(), ShouldEqual, "aaa")

			Convey("inout restore by snap", func() {
				in.Restore(snap)
				So(in.GetText(), ShouldEqual, "aaa")
			})
		})

	})
}
