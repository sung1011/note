package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_FlyWeight(t *testing.T) {
	Convey("", t, func() {
		Convey("generate different chessboard A and B", func() {
			cb_A := NewChessBoard()
			cb_A.Move(1, 2, 3)

			cb_B := NewChessBoard()
			cb_B.Move(2, 8, 9)
			Convey("but same unit", func() {
				So(cb_A.chessPieces[1].Unit, ShouldEqual, cb_B.chessPieces[1].Unit)
				So(cb_A.chessPieces[2].Unit, ShouldEqual, cb_B.chessPieces[2].Unit)
			})
		})

	})
}
