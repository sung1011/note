package main

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// TODO ascii

// TODO utf8

func Test_fmt(t *testing.T) {
	Convey("", t, func() {

		tmpByteSlice := []byte("go开发")

		type sample struct {
			Name string
			Age  int
		}

		var fmtTests = []struct {
			fmt string
			val interface{}
			out string
		}{

			// integers %d
			{"%d", ^uint8(0), "255"},
			{"%d", ^uint16(0), "65535"},
			{"%d", ^uint32(0), "4294967295"},
			{"%d", ^uint64(0), "18446744073709551615"},
			// integers %b
			{"%b", 7, "111"},

			// characters 字节码 %c
			{"%c", 0x672c, "本"},
			{"%c", 0x4E2D, "中"},
			{"%c", '日', "日"},
			{"%c", -1, "�"}, // not valid

			// escaped characters %q
			{"%q", 0x4E2D, "'中'"}, // 转子转unicode, 该值对应的单引号括起来
			{"%q", '"', `'"'`},
			{"%q", uint(0), `'\x00'`},
			{"%q", 'x', `'x'`},
			{"%q", '\n', `'\n'`},

			// width %d
			{"|%5d|", 123, "|  123|"},        // 右缩进
			{"|%-5d|", 123, "|123  |"},       // 左缩进
			{"|%05d|", 123, "|00123|"},       // 补位
			{"|%05d|", 1234567, "|1234567|"}, // 补位, 但无需补
			{"|%.5d|", 123, "|00123|"},       // 同%05d
			{"|%.5d|", 1234567, "|1234567|"}, // 同%05d

			// width %s
			{"|%5s|", "abc", "|  abc|"},
			{"|%-5s|", "abc", "|abc  |"},
			{"|%05s|", "abc", "|00abc|"},
			{"|%5s|", "abcdefg", "|abcdefg|"}, //
			{"|%.5s|", "abcdefg", "|abcde|"},  // 截取, 最多5个字符
			{"|%.5s|", "abc", "|abc|"},        // 截取, 最多5个字符

			// 小数点 %f
			{"%+.3f", -1.0, "-1.000"},

			// %e 科学计数法
			{"% .3e", 1.0, " 1.000e+00"},
			{"% .3E", -1.0, "-1.000E+00"},

			// %g 用最少的数字表示
			{"%+.3g", 0.0, "+0"},
			{"%+.3g", 1.0, "+1"},
			{"%+.3g", -1.0, "-1"},
			{"% .3g", -1.0, "-1"},
			{"% .3g", 1.0, " 1"},

			// string
			{"%s", []byte("go开发"), "go开发"}, // toString
			{"%x", "xyz", "78797a"},
			{"% x", "xyz", "78 79 7a"}, // 带缩进
			{"%#x", "xyz", "0x78797a"},
			{"%# x", "xyz", "0x78 0x79 0x7a"},
			{"%q", "go开发", `"go开发"`}, // 双引号围绕的字符串, 由Go语法安全地转义

			// struct
			{"%v", sample{"tickles", 32}, "{tickles 32}"},                           // 值
			{"%+v", sample{"tickles", 32}, "{Name:tickles Age:32}"},                 // 键值
			{"%#v", sample{"tickles", 32}, "main.sample{Name:\"tickles\", Age:32}"}, // 类型, 键值
			{"%T", sample{"tickles", 32}, "main.sample"},                            // 类型

			// bool
			{"%t", true, "true"}, // bool

			// 整数
			{"%v%%", 32, "32%"},      // 输出%本体
			{"%c", 0x4E2D, "中"},      // 数字转unicode码
			{"%q", 0x4E2D, "'中'"},    // 转子转unicode, 该值对应的单引号括起来
			{"%U", 0x4E2D, "U+4E2D"}, //  数字转Unicode字符串 \u4e2d
			{"%b", 26, "11010"},      // bin 2
			{"%d", 0x12, "18"},       // dec 10
			{"%o", 20, "24"},         // octed 8
			{"%x", 14, "e"},          // hex 16 a~f
			{"%X", 14, "E"},          // hex 16 A~F
			{"%d", 12345, "12345"},   // int

			// 浮点数
			{"%b", 10.45, "5882827013252710p-49"}, // 无小数部分, 二进制的科学计数法
			{"%e", 10.45, "1.045000e+01"},         // 科学计数法(小写)
			{"%E", 10.45, "1.045000E+01"},         // 科学计数法(大写)
			{"%f", 10.45, "10.450000"},            // 有小数, 无指数
			{"%F", 10.45, "10.450000"},            // 同%f
			{"%g", 10.45, "10.45"},                // 根据实际情况采用%e或%f格式(以获得更简洁、准确的输出)
			{"%G", 10.45, "10.45"},                // 根据实际情况采用%E或%F格式(以获得更简洁、准确的输出)

			// []byte
			{"%p", tmpByteSlice, fmt.Sprintf("%p", tmpByteSlice)},   // 切片第一个元素的指针 带0x前缀
			{"%#p", tmpByteSlice, fmt.Sprintf("%#p", tmpByteSlice)}, // 切片第一个元素的指针

		}
		for _, row := range fmtTests {
			t.Logf("%v %v %v", row.fmt, row.val, row.out)
			So(fmt.Sprintf(row.fmt, row.val), ShouldEqual, row.out)
		}
	})
}
