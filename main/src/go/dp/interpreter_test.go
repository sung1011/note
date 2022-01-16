package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Interpreter(t *testing.T) {
	Convey("interpret something", t, func() {
		stats := map[string]float64{
			"a": 1,
			"b": 2,
			"c": 3,
		}
		Convey("true", func() {
			ar, _ := NewAlertRule("a > 0 && b < 10 && c < 5")
			So(ar.Interpret(stats), ShouldEqual, true)
		})
		Convey("false", func() {
			ar, _ := NewAlertRule("a > 1 && b < 10 && c < 5")
			So(ar.Interpret(stats), ShouldEqual, false)
		})
		Convey("err", func() {
			_, err := NewAlertRule("a == 1 && b < 10 && c < 5")
			// t.Log(err)
			So(err, ShouldBeError)
		})

	})
}
