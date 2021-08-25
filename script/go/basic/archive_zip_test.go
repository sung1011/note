package main

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Archive_Zip(t *testing.T) {
	Convey("", t, func() {
		files := map[string]string{
			"context_test.go": "This archive contains some text files.",
			"defer_test.go":   "Gopher names:\nGeorge\nGeoffrey\nGonzo",
		}
		// files := []struct {
		// 	Name, Body string
		// }{
		// 	{"context_test.go", "This archive contains some text files."},
		// 	{"defer_test.go", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		// 	// {"todo.txt", "Get animal handling license."},
		// }
		Convey("write", func() {

			f, _ := os.Create("archive_zip_test.gz")
			defer os.Remove("archive_zip_test.gz")
			zw := zip.NewWriter(f)
			for name, body := range files {
				w, _ := zw.Create(name)
				w.Write([]byte(body))
			}
			zw.Close() // flush

			Convey("read", func() {
				zr, err := zip.OpenReader(f.Name())
				So(err, ShouldBeNil)
				for _, zrf := range zr.File {
					var buf bytes.Buffer
					rc, _ := zrf.Open()
					io.Copy(&buf, rc)
					So(buf.String(), ShouldEqual, files[zrf.FileInfo().Name()])
				}
			})
		})
	})
}
