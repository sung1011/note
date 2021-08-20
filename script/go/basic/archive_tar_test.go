package main

import (
	"archive/tar"
	"bytes"
	"io"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Archive_Tar(t *testing.T) {
	Convey("", t, func() {
		files := []struct {
			Name, Body string
		}{
			{"readme.txt", "This archive contains some text files."},
			{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
			{"todo.txt", "Get animal handling license."},
		}
		Convey("write", func() {
			// write
			var buf bytes.Buffer
			tw := tar.NewWriter(&buf)
			for _, file := range files {
				err := tw.WriteHeader(&tar.Header{
					Name: file.Name,
					Mode: 0666,
					Size: int64(len(file.Body)),
				})
				So(err, ShouldBeNil)

				_, err = tw.Write([]byte(file.Body))
				So(err, ShouldBeNil)
			}
			defer tw.Close()

			// read
			Convey("read", func() {
				tr := tar.NewReader(&buf)
				for {
					header, err := tr.Next()
					if err == io.EOF {
						break
					}
					So(err, ShouldBeNil)

					So(header.Name, ShouldBeIn, []string{"readme.txt", "gopher.txt", "todo.txt"})

					b := make([]byte, header.Size)
					_, err = tr.Read(b)
					// t.Log(string(b))
					So(err, ShouldEqual, io.EOF)
					So(string(b), ShouldBeIn, []string{"This archive contains some text files.", "Gopher names:\nGeorge\nGeoffrey\nGonzo", "Get animal handling license."})
				}
			})
		})
	})
}
