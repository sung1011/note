package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Visitor(t *testing.T) {
	Convey("", t, func() {
		compressor := &Compressor{}
		Convey("pdf", func() {
			pdf, _ := NewResourceFile("./xx.pdf")
			rsPdf, _ := pdf.Accept(compressor)
			So(rsPdf, ShouldEqual, "pdf file")

		})
		Convey("ppt", func() {
			ppt, _ := NewResourceFile("./xx.ppt")
			rsPPT, _ := ppt.Accept(compressor)
			So(rsPPT, ShouldEqual, "ppt file")
		})

		Convey("xxx", func() {
			_, err := NewResourceFile("./xx.xyz")
			So(err, ShouldBeError)
		})
	})
}
