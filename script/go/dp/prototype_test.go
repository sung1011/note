package main

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Prototype(t *testing.T) {
	t1, _ := time.Parse("2006", "1999")
	wm := WordMap{
		"foo": &Word{s: "foo", visit: 11, UpdAt: &t1},
		"bar": &Word{s: "bar", visit: 22, UpdAt: &t1},
		"baz": &Word{s: "baz", visit: 33, UpdAt: &t1},
	}

	newBarVisit := 29
	newBazVisit := 39
	t2, _ := time.Parse("2006", "2021")
	ws := []*Word{
		{s: "bar", visit: newBarVisit, UpdAt: &t2},
		{s: "baz", visit: newBazVisit},
	}
	nwm := wm.Clone(ws)
	t.Log(nwm)

	Convey("", t, func() {
		So(nwm["bar"].visit, ShouldEqual, newBarVisit)
		So(nwm["baz"].visit, ShouldEqual, newBazVisit)
	})
}
