package main

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// bufio比buffer复杂一点, 多了个内存缓冲区

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

		w1 := bufio.NewWriter(bf) // 新建缓冲区
		w1.WriteString("a")       // 写入w1的缓冲
		w1.WriteString("b")
		w1.WriteString("c")
		t.Log(w1.Size())                  // 4096; defaultBufSize=4096
		So(w1.Buffered(), ShouldEqual, 3) // 没有flush, 数据还在w1的缓冲区
		So(bf.Len(), ShouldEqual, 0)      // 没有flush, 所以尚未写入bf

		newbf := new(bytes.Buffer)
		w1.Reset(newbf)
		So(w1.Buffered(), ShouldEqual, 0) // reset清空缓冲区, 并装载新的bf

		w2 := bufio.NewWriterSize(bf, w1.Size()*2) // 新建缓冲区(by 指定size)
		w2.WriteString("x")                        // 写入w的缓冲
		w2.WriteString("y")
		So(w2.Buffered(), ShouldEqual, 2) // 没有flush, 数据还在w1的缓冲区, 新建了, 所以是2而非3+2
		So(bf.Len(), ShouldEqual, 0)      // 尚未写入bf
		w2.Flush()                        // 将w的缓冲写入bf

		So(w2.Buffered(), ShouldEqual, 0) // w2的缓冲区已经被清空
		So(bf.Len(), ShouldEqual, 2)      // 确实已经写入bf了
	})
}
