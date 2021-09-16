package main

import (
	"crypto/md5"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_md5(t *testing.T) {
	Convey("", t, func() {
		// 方式1
		m := md5.New()
		m.Write([]byte("123456"))
		rs1 := m.Sum(nil)

		// 方式2
		rs2 := md5.Sum([]byte("123456"))

		hexRs1 := fmt.Sprintf("%x", rs1)
		hexRs2 := fmt.Sprintf("%x", rs2)

		So(hexRs1, ShouldEqual, hexRs2) // e10adc3949ba59abbe56e057f20f883e
	})
}
