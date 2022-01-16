package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Responsibility_chain(t *testing.T) {
	Convey("", t, func() {
		chain := &SensitiveWordFilterChain{}
		Convey("just filt ad", func() {
			chain.AddFilter(&AdSensitiveWordFilter{})
			So(chain.Filter("test"), ShouldEqual, false)
			Convey("filt ad and political", func() {
				chain.AddFilter(&PoliticalWordFilter{})
				So(chain.Filter("test"), ShouldEqual, true)
			})
		})

	})
}
