package main

import (
	"bytes"
	"io"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// Buffer结构很简单, 就是个slice (buf)
//	buf      []byte // contents are the bytes buf[off : len(buf)]
//  off      int    // read at &buf[off], write at &buf[len(buf)]
//  lastRead readOp // last read operation, so that Unread* can work correctly.

var buffer_txt string = "github.com\nsmartystreets\ngoconvey\nconvey"

func Test_Buffer_Read(t *testing.T) {
	Convey("", t, func() {
		var err error
		var i int
		bf := bytes.NewBuffer([]byte(txt))
		for err == nil {
			_, err := bf.ReadBytes('\n')
			if err == io.EOF {
				break
			}
			So(err, ShouldBeNil)
			switch i {
			case 0:
				So(bf.Len(), ShouldEqual, 29)               // 算上\n分隔符; 实际 smartystreets\ngoconvey\nconvey
				So(bf.Cap(), ShouldEqual, cap([]byte(txt))) // buffer中buf切片的cap = new时的cap
			case 1:
				So(bf.Len(), ShouldEqual, 15)
			case 2:
				So(bf.Len(), ShouldEqual, 6)
			case 3:
				So(bf.Len(), ShouldEqual, 0)
			}
			i++
		}
	})
}

func Test_Buffer_Write(t *testing.T) {
	Convey("", t, func() {
		Convey("new()", func() {
			bf := new(bytes.Buffer) // writer
			bf.WriteString("a")
			bf.WriteString("b")
			bf.WriteString("c")
			So(bf.Cap(), ShouldBeGreaterThan, 3) // new()出的cap大于3个
			So(bf.Len(), ShouldEqual, 3)         // 写入3个
		})

		Convey("NewBuffer()", func() {
			bf := bytes.NewBuffer([]byte(""))
			bf.WriteString("a")
			bf.WriteString("b")
			bf.WriteString("c")
			So(bf.Cap(), ShouldEqual, 3) // NewBuffer的参数值决定了cap的大小
			So(bf.Len(), ShouldEqual, 3) // 写入3个
		})

		Convey("reset", func() {
			bf := bytes.NewBuffer([]byte(""))
			bf.WriteString("a")
			bf.WriteString("b")
			bf.Reset() // 清空buf
			bf.WriteString("c")
			So(bf.Len(), ShouldEqual, 1) // 只有一个c
			So(bf.Cap(), ShouldEqual, 3) // cap已经被扩展为3了
		})

		Convey("grow cap", func() {
			bf := bytes.NewBuffer([]byte(""))
			bf.Grow(100)
			So(bf.Cap(), ShouldEqual, 100)
		})

		Convey("truncate", func() {
			bf := bytes.NewBuffer([]byte(buffer_txt))
			bs := bf.Next(3) // offset偏移量+=3
			So(string(bs), ShouldEqual, "git")
			So(string(bf.Bytes()), ShouldEqual, "hub.com\nsmartystreets\ngoconvey\nconvey") // buf="hub.com......."
			bf.Truncate(3)                                                                  // 从头向后[:off+n]剪切出3个字符, 其他都丢弃
			So(string(bf.Bytes()), ShouldEqual, "hub")                                      // buf="hub"

			Convey("next", func() {
				bf.Next(1)                                // offset偏移量+=1
				So(string(bf.Bytes()), ShouldEqual, "ub") // buf="ub"
			})
		})
	})
}
