package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Archive_Zip(t *testing.T) {
	Convey("", t, func() {
		files := []struct {
			Name, Body string
		}{
			{"readme.txt", "This archive contains some text files."},
			{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
			{"todo.txt", "Get animal handling license."},
		}
		Convey("write", func() {

			Convey("read", func() {

			})
		})
	})
}
