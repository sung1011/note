package main

import (
	"errors"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Errors(t *testing.T) {
	Convey("", t, func() {
		Convey("errors.Is()", func() {
			So(errors.Is(errors.New("lala"), errors.New("lala")), ShouldBeFalse)
		})

		Convey("errors.As()", func() {
			if _, err := os.Open("non-existing"); err != nil {
				var pathError *os.PathError
				So(errors.As(err, &pathError), ShouldBeTrue)

			}
		})

		Convey("assert", func() {
			_, err := os.Open("non-existing")
			_, ok := err.(*os.PathError)
			So(ok, ShouldBeTrue)
		})
	})
}
