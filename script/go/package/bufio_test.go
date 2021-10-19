package main

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var txt string = "github.com\nsmartystreets\ngoconvey\nconvey"

func Test_Bufio_Read(t *testing.T) {
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
				So(rd.Buffered(), ShouldEqual, 29)      // 算上\n分隔符; 实际 smartystreets\ngoconvey\nconvey
				l1, _ := rd.Peek(6)                     // smarty 拿到顶部数据, 但不取出
				l2, _ := rd.Peek(6)                     // smarty 拿到顶部数据, 但不取出
				So(string(l1), ShouldEqual, string(l2)) // smarty
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

func Test_Bufio_Write(t *testing.T) {
	Convey("", t, func() {
		bf := new(bytes.Buffer) // writer
		w := bufio.NewWriter(bf)
		w.WriteString("a") // 写入w的缓冲
		w.WriteString("b")
		w.WriteString("c")
		So(w.Buffered(), ShouldEqual, 3)
		w.Flush() // 将w的缓冲写入writer

		So(w.Buffered(), ShouldEqual, 0) // w的缓冲区已经被清空
		So(bf.Len(), ShouldEqual, 3)     // 确实已经写入writer了
	})
}
