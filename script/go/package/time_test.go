package main

import (
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Time_TZ(t *testing.T) {
	Convey("Timezone", t, func() {
		now := time.Now()

		So(now.Location().String(), ShouldEqual, "Local") // 默认Local, 实际是/etc/localtime的值

		loc, _ := time.LoadLocation("Etc/GMT+4")
		So(now.In(loc).Location().String(), ShouldEqual, "Etc/GMT+4")

		So(now.UTC().Location().String(), ShouldEqual, "UTC")
	})
}

func Test_Time_Format(t *testing.T) {
	Convey("Format", t, func() {
		now := time.Now()
		s1 := now.Format("2006-01-02 15:04:05")
		s2 := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
		So(s1, ShouldEqual, s2)
	})
}

func Test_Time_Ts(t *testing.T) {
	Convey("Ts", t, func() {
		now := time.Now()
		timestamp := now.Unix() // time.Time -> ts

		t := time.Unix(timestamp, 0)                                                        // ts -> time.Time
		So(now.Format("2006-01-02 15:04:05"), ShouldEqual, t.Format("2006-01-02 15:04:05")) // 相同
	})
}

func Test_Time_Func(t *testing.T) {
	Convey("Func", t, func() {
		now := time.Now()
		later := now.Add(1 * time.Hour)
		So(later.Unix(), ShouldEqual, now.Unix()+3600)
		So(later.After(now), ShouldBeTrue)
		So(now.Before(later), ShouldBeTrue)
		So(later.Sub(now), ShouldEqual, 1*time.Hour)
	})
}

func Test_Time_Parse(t *testing.T) {
	Convey("Parse", t, func() {
		t, _ := time.Parse("2006-01-02 15:04:05", "2021-10-11 04:00:00") // str -> time.Time
		So(t.Format("2006-01-02 15:04:05"), ShouldEqual, "2021-10-11 04:00:00")
	})
}
