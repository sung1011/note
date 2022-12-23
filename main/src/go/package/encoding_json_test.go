package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Marshal(t *testing.T) {
	Convey("Marshal() struct{} -> string", t, func() {
		type ColorGroup struct {
			ID     int
			Name   string
			Colors []string
		}
		group := ColorGroup{
			ID:     1,
			Name:   "Reds",
			Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
		}
		b, err := json.Marshal(group)
		So(err, ShouldBeNil)
		So(string(b), ShouldEqual, `{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}`)
	})
}

func Test_Unmarshal(t *testing.T) {
	Convey("Unmarshal() string -> struct{}", t, func() {
		var jsonBlob = []byte(
			`[
					{"Name": "Platypus", "Order": "Monotremata"},
					{"Name": "Quoll",    "Order": "Dasyuromorphia"}
				]`)
		type Animal struct {
			Name  string
			Order string
		}
		var animals []Animal

		err := json.Unmarshal(jsonBlob, &animals)
		So(err, ShouldBeNil)
		So("Quoll", ShouldEqual, animals[1].Name)
		// t.Logf("%+v", animals) // [{Name:Platypus Order:Monotremata} {Name:Quoll Order:Dasyuromorphia}]
	})
}

func Test_Marshal_With_Tag(t *testing.T) {
	Convey("Marshal() with tag struct{} -> string", t, func() {
		type ColorGroup struct {
			ID     int      `json:"id,string"`       // 转类型; int转str
			Name   string   `json:"color_name"`      // 转名字; name转color_name
			Colors []string `json:"-"`               // 忽略; 不转出该字段
			Label  int      `json:"color,omitempty"` // 省略空; 值为空值则不转出该字段
		}
		group := ColorGroup{
			ID:     1,
			Name:   "Reds",
			Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
			Label:  0,
		}
		b, err := json.Marshal(group)
		So(err, ShouldBeNil)
		So(string(b), ShouldEqual, `{"id":"1","color_name":"Reds"}`)
	})
}

func Test_Decoder(t *testing.T) {
	Convey("Decoder", t, func() {
		const jsonStream = `
	{"Name": "Ed", "Text": "Knock knock."}
	{"Name": "Sam", "Text": "Who's there?"}
	{"Name": "Ed", "Text": "Go fmt."}
	{"Name": "Sam", "Text": "Go fmt who?"}
	{"Name": "Ed", "Text": "Go fmt yourself!"}
`
		type Message struct {
			Name, Text string
		}
		dec := json.NewDecoder(strings.NewReader(jsonStream))
		var m Message
		for {
			if err := dec.Decode(&m); err == io.EOF {
				break
			} else if err != nil {
				t.Fatal(err)
			}
			So(m.Name, ShouldBeIn, []string{"Ed", "Sam"})
			So(m.Text, ShouldBeIn, []string{"Knock knock.", "Who's there?", "Go fmt.", "Go fmt who?", "Go fmt yourself!"})
			t.Log(m)
		}
	})
}

func Test_Decoder_Token(t *testing.T) {
	Convey("Debug decoder", t, func() {
		const jsonStream = `
	{"Message": "Hello", "Array": [1, 2, 3], "Null": null, "Number": 1.234}
`
		dec := json.NewDecoder(strings.NewReader(jsonStream))
		for {
			tk, err := dec.Token()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			t.Logf("%T: %v", tk, tk)
			if dec.More() {
				t.Logf(" (more)")
			}
		}
		// Output:
		// json.Delim: { (more)
		// string: Message (more)
		// string: Hello (more)
		// string: Array (more)
		// json.Delim: [ (more)
		// float64: 1 (more)
		// float64: 2 (more)
		// float64: 3
		// json.Delim: ] (more)
		// string: Null (more)
		// <nil>: <nil> (more)
		// string: Number (more)
		// float64: 1.234
		// json.Delim: }
	})
}

func Test_RawMessage_unmarshal(t *testing.T) {
	Convey("json.RawMessage", t, func() {
		type Color struct {
			Space string
			Point json.RawMessage // delay parsing until we know the color space
		}
		type RGB struct {
			R uint8
			G uint8
			B uint8
		}
		type YCbCr struct {
			Y  uint8
			Cb int8
			Cr int8
		}

		var j = []byte(`[
	{"Space": "YCbCr", "Point": {"Y": 255, "Cb": 0, "Cr": -10}},
	{"Space": "RGB",   "Point": {"R": 98, "G": 218, "B": 255}}
]`)
		var colors []Color
		err := json.Unmarshal(j, &colors)
		if err != nil {
			log.Fatalln("error:", err)
		}

		for _, c := range colors {
			var dst interface{}
			switch c.Space {
			case "RGB":
				dst = new(RGB)
			case "YCbCr":
				dst = new(YCbCr)
			}
			err := json.Unmarshal(c.Point, dst)
			if err != nil {
				log.Fatalln("error:", err)
			}
			// fmt.Printf("%+v", dst)
			// fmt.Println(c.Space, dst)
			So(c.Space, ShouldBeIn, []string{"YCbCr", "RGB"})
			So(fmt.Sprintf("%+v", dst), ShouldBeIn, []string{"&{Y:255 Cb:0 Cr:-10}", "&{R:98 G:218 B:255}"})
		}
	})
}

func Test_Indent(t *testing.T) {
	Convey("Indent", t, func() {
		type Road struct {
			Name   string
			Number int
		}
		roads := []Road{
			{"Diamond Fork", 29},
			{"Sheep Creek", 51},
		}

		b, err := json.Marshal(roads)
		if err != nil {
			log.Fatal(err)
		}

		var out bytes.Buffer
		json.Indent(&out, b, "=", "\t")

		// Output:
		output := `[
=	{
=		"Name": "Diamond Fork",
=		"Number": 29
=	},
=	{
=		"Name": "Sheep Creek",
=		"Number": 51
=	}
=]`
		So(out.String(), ShouldEqual, output)
	})
}

func Test_MarshalIndent(t *testing.T) {
	Convey("marshal indent", t, func() {
		data := map[string]int{
			"a": 1,
			"b": 2,
		}

		json, err := json.MarshalIndent(data, "<prefix>", "<indent>")
		if err != nil {
			log.Fatal(err)
		}

		output := `{
<prefix><indent>"a": 1,
<prefix><indent>"b": 2
<prefix>}`
		So(string(json), ShouldEqual, output)
	})
}

func Test_Valid(t *testing.T) {
	Convey("", t, func() {
		goodJSON := `{"example": 1}`
		So(json.Valid([]byte(goodJSON)), ShouldBeTrue)

		badJSON := `{"example":2:]}}`
		So(json.Valid([]byte(badJSON)), ShouldBeFalse)
	})
}

func Test_HTMLEscape(t *testing.T) {
	Convey("", t, func() {
		var out bytes.Buffer
		json.HTMLEscape(&out, []byte(`{"Name":"<b>HTML content</b>"}`))
		So(out.String(), ShouldEqual, `{"Name":"\u003cb\u003eHTML content\u003c/b\u003e"}`)
	})
}
