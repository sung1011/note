package main

import (
	"bufio"
	"io"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var txt string = "github.com\nsmartystreets\ngoconvey\nconvey"

func Test_Bufio(t *testing.T) {
	Convey("", t, func() {
		var err error
		var i int
		rd := bufio.NewReader(strings.NewReader(txt))
		for err == nil {
			bs, err := rd.ReadBytes('\n')
			t.Log(string(bs), rd.Buffered())
			if err == io.EOF {
				break
			}
			So(err, ShouldBeNil)
			// buffered 缓冲中的字节数
			switch i {
			case 0:
				So(rd.Buffered(), ShouldEqual, 29)
				l1, _ := rd.Peek(6) // smarty
				l2, _ := rd.Peek(6) // smarty
				So(string(l1), ShouldEqual, string(l2))
			case 1:
				So(rd.Buffered(), ShouldEqual, 15)
			case 2:
				So(rd.Buffered(), ShouldEqual, 6)
			case 3:
				So(rd.Buffered(), ShouldEqual, 0)
			}
			i++
		}
	})
}
