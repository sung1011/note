package main

import (
	"archive/zip"
	"io"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Archive_Zip(t *testing.T) {
	Convey("", t, func() {
		files := []struct {
			Name, Body string
		}{
			{"context_test.go", "This archive contains some text files."},
			{"defer_test.go", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
			// {"todo.txt", "Get animal handling license."},
		}
		Convey("write", func() {

			f, _ := os.Create("archive_zip_test.gz")
			zw := zip.NewWriter(f)
			for _, file := range files {

				zw.Create(file.Name)
				defer zw.Close()
			}

			Convey("read", func() {
				zr, err := zip.OpenReader(f.Name())
				So(err, ShouldBeNil)
				for _, zrf := range zr.File {
					rc, _ := zrf.Open()
					io.Copy(os.Stdout, rc)
				}
			})
		})
	})
}
