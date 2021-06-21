package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Observer(t *testing.T) {
	Convey("", t, func() {
		n := &NewsOffice{}

		Convey("news subscribe", func() {
			a := &CustomerA{}
			n.addCustomer(a)

			b := &CustomerB{}
			n.addCustomer(b)

			Convey("news public and notify", func() {
				n.newspaperCome()

				So(a.recv, ShouldEqual, true)
				So(b.recv, ShouldEqual, true)
			})
		})

	})
}
