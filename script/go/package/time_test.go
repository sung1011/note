package main

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Time(t *testing.T) {
	Convey("Timezone", t, func() {
		now := time.Now()

		So(now.Location().String(), ShouldEqual, "Local") // 默认Local, 实际是/etc/localtime的值

		loc, _ := time.LoadLocation("Etc/GMT+4")
		So(now.In(loc).Location().String(), ShouldEqual, "Etc/GMT+4")

		So(now.UTC().Location().String(), ShouldEqual, "UTC")
	})
}
