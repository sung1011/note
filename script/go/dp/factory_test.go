package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Simple(t *testing.T) {
	Convey("new parser", t, func() {
		So(NewIParser("json"), ShouldEqual, &jsonParser{})
		So(NewIParser("yaml"), ShouldEqual, &yamlParser{})
		So(NewIParser("___"), ShouldEqual, nil)
	})
}

func Test_Method(t *testing.T) {
	Convey("new parser factory", t, func() {
		So(NewParserFactory("json"), ShouldEqual, &jsonParserFactory{})
		So(NewParserFactory("yaml"), ShouldEqual, &yamlParserFactory{})
		So(NewParserFactory("___"), ShouldEqual, nil)
	})
}

func Test_Abstract_GenTextParser(t *testing.T) {
	Convey("new parser factory", t, func() {
		spf := &sqlParserFactory{}
		So(spf.GenIParserToText(), ShouldEqual, &sqlParserText{})
		So(spf.GenIParserToPic(), ShouldEqual, &sqlParserPic{})
	})
}
