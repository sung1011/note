package main

import (
	"archive/zip"
	"bufio"
	"compress/gzip"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Create(t *testing.T) {
	Convey("create empty file", t, func() {

		o := New(tmp_file)

		Convey("file not exist", func() {
			So(o.Exist(), ShouldBeFalse)
			Convey("create()", func() {
				_, err := o.Create()
				defer o.Remove()
				So(err, ShouldBeNil)
				Convey("file exist", func() {
					So(o.Exist(), ShouldBeTrue)
				})
			})
		})
	})
}

func Test_Truncate(t *testing.T) {
	Convey("truncate file content", t, func() {

		o := New(tmp_file)

		content := "abcdefg"
		Convey("init file -> "+content, func() {
			err := o.WriteSimple([]byte(content))
			defer o.Remove()
			So(err, ShouldBeNil)
			Convey(fmt.Sprintf("truncate(3) %s -> abc", content), func() {
				err = o.Truncate(3)
				So(err, ShouldBeNil)
				bs, err := o.ReadSimple()
				So(err, ShouldBeNil)
				So(bs, ShouldResemble, []byte("abc"))

			})
			Convey(fmt.Sprintf("truncate(10) %s -> %s<nil><nil><nil>", content, content), func() {
				err = o.Truncate(10)
				So(err, ShouldBeNil)
				bs, err := o.ReadSimple()
				So(err, ShouldBeNil)
				// t.Log(bs, len(bs), string(bs))
				rs := append([]byte(content), 0, 0, 0)
				So(bs, ShouldResemble, rs)
			})

			Convey(fmt.Sprintf("truncate(0) %s -> ", content), func() {
				err = o.Truncate(0)
				So(err, ShouldBeNil)
				bs, err := o.ReadSimple()
				So(err, ShouldBeNil)
				So(bs, ShouldBeEmpty)
			})
		})
	})
}

func Test_Permission(t *testing.T) {
	Convey("check permission", t, func() {

		o := New(tmp_file)

		Convey("create file", func() {
			_, err := o.Open(os.O_CREATE|os.O_RDONLY|os.O_TRUNC, 0666)
			So(err, ShouldBeNil)

			Convey("chmod rw -> r", func() {
				err := os.Chmod(tmp_file, 0444)
				So(err, ShouldBeNil)
				Convey("check permission write", func() {
					_, err := o.Open(os.O_WRONLY, 0666)
					So(os.IsPermission(err), ShouldBeTrue)

					Convey("chown", func() {
						err := os.Chown(tmp_file, os.Getuid(), os.Getgid())
						// TODO 如何获取文件的own, group?
						So(err, ShouldBeNil)
					})
				})
			})
		})
		defer o.Remove()
	})
}

func Test_SymLink(t *testing.T) {
	Convey("", t, func() {
		o := New(tmp_file)

		linkPath := "tmp_symlink"

		Convey("create file", func() {
			_, err := o.Open(os.O_CREATE|os.O_RDONLY|os.O_TRUNC, 0666)
			So(err, ShouldBeNil)

			Convey("sym link", func() {
				err := os.Symlink(o.path, linkPath)
				So(err, ShouldBeNil)

				fileInfo, err := os.Lstat(linkPath)
				So(err, ShouldBeNil)

				So(fileInfo.Name(), ShouldEqual, linkPath)
				So(fileInfo.Mode().IsRegular(), ShouldBeFalse)
				So(fileInfo.Mode()&os.ModeSymlink == os.ModeSymlink, ShouldBeTrue)
				// 改变软连的拥有者, 不会影原始文件
			})
		})
		o.Remove()
		defer os.Remove(linkPath)
	})
}

func Test_Copy(t *testing.T) {
	Convey("", t, func() {
		o := New(tmp_file)

		targetPath := "tmp_copy"

		Convey("create file", func() {
			o.WriteSimple([]byte(`foo
bar
baz
qux
...
`))
			Convey("copy", func() {
				o.Copy(targetPath)

				bs, err := ioutil.ReadFile(targetPath)
				So(err, ShouldBeNil)
				oldBs, err := o.ReadSimple()
				So(err, ShouldBeNil)
				So(len(bs), ShouldEqual, len(oldBs))

				defer o.Remove()
				defer os.Remove(targetPath)
			})
		})
	})
}

func Test_Seek(t *testing.T) {
	Convey("", t, func() {
		o := New(tmp_file)

		Convey("create file", func() {
			o.WriteSimple([]byte(`foo
bar
baz
qux
`))
			defer o.Remove()
			Convey("move cursor", func() {
				f, err := os.OpenFile(o.path, os.O_WRONLY, 0666)
				defer f.Close()
				So(err, ShouldBeNil)

				_, err = f.Seek(5, 0)
				So(err, ShouldBeNil)

				Convey("write content", func() {
					_, err = f.Write([]byte("^_^"))
					So(err, ShouldBeNil)

					bs, err := ioutil.ReadFile(tmp_file)
					So(err, ShouldBeNil)
					So(bs, ShouldResemble, []byte(`foo
b^_^baz
qux
`))
				})
			})
		})
	})
}

func Test_Write(t *testing.T) {
	Convey("", t, func() {

		s := "foo\nbar\nbaz"
		l := len(s) // 11

		Convey("wirte file", func() {
			f, err := os.OpenFile(tmp_file, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
			defer os.Remove(tmp_file)
			defer f.Close()
			So(err, ShouldBeNil)
			n, err := f.Write([]byte(s))
			So(err, ShouldBeNil)
			So(n, ShouldEqual, l)
		})

		Convey("write file sample", func() {
			err := ioutil.WriteFile(tmp_file, []byte(s), 0666)
			defer os.Remove(tmp_file)
			So(err, ShouldBeNil)
			bs, err := ioutil.ReadFile(tmp_file)
			So(err, ShouldBeNil)
			So(len(bs), ShouldEqual, l)
		})

		// 在内存中攒一波, 一起刷入硬盘, 减少磁盘IO, 提升效率
		Convey("write file buffer", func() {
			f, err := os.OpenFile(tmp_file, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
			defer os.Remove(tmp_file)
			defer f.Close()
			So(err, ShouldBeNil)

			w := bufio.NewWriter(f)

			w.Write([]byte("foo\n"))

			w.WriteString("bar\n")

			w.WriteByte(98)  // b
			w.WriteByte(97)  // a
			w.WriteByte(122) // z

			n := w.Buffered()
			So(n, ShouldEqual, l)

			w.Flush()

			bs, err := ioutil.ReadFile(tmp_file)
			So(err, ShouldBeNil)

			So(len(bs), ShouldEqual, l)
		})
	})
}

func Test_Read(t *testing.T) {
	Convey("", t, func() {
		s := "foo\nbar\nbaz"
		l := len(s)
		err := ioutil.WriteFile(tmp_file, []byte(s), 0666)
		So(err, ShouldBeNil)
		defer os.Remove(tmp_file)

		Convey("read file", func() {
			// TODO 读取最多n个字节
			// TODO 读取正好n个字节
			// TODO 读取至少n个字节
			// 读取全部字节
			f, err := os.Open(tmp_file)
			So(err, ShouldBeNil)
			bs, err := ioutil.ReadAll(f)
			So(err, ShouldBeNil)
			So(len(bs), ShouldEqual, l)
		})
		Convey("read file sample", func() {
			bs, err := ioutil.ReadFile(tmp_file)
			So(err, ShouldBeNil)
			So(len(bs), ShouldEqual, l)
		})
		Convey("read file buffer", func() {
			f, err := os.Open(tmp_file)
			So(err, ShouldBeNil)
			r := bufio.NewReader(f)

			Convey("buffer-reader peek, 读取 但不移动指针", func() {
				bs := make([]byte, 2)
				bs, err = r.Peek(2)
				So(err, ShouldBeNil)
				So(bs, ShouldResemble, []byte("fo"))
			})

			Convey("buffer-reader read, 读取 并移动指针", func() {
				bs := make([]byte, 2)
				_, err = r.Read(bs)
				So(err, ShouldBeNil)
				So(bs, ShouldResemble, []byte("fo"))
			})

			Convey("buffer-readbyte, 读取一字节, 不成功则err", func() {
				b, err := r.ReadByte()
				So(err, ShouldBeNil)
				So(b, ShouldEqual, byte('f'))
			})

			Convey("buffer-readbytes, 读取 至分隔符(包含分隔符)", func() {
				bs, err := r.ReadBytes('\n')
				So(err, ShouldBeNil)
				So(bs, ShouldResemble, []byte("foo\n"))
			})
		})
	})
}

func Test_Scanner(t *testing.T) {
	Convey("", t, func() {
		s := "foo\nbar\nbaz"
		// l := len(s)
		err := ioutil.WriteFile(tmp_file, []byte(s), 0666)
		So(err, ShouldBeNil)
		defer os.Remove(tmp_file)

		Convey("scanner", func() {
			f, err := os.Open(tmp_file)
			So(err, ShouldBeNil)
			scanner := bufio.NewScanner(f)

			// bufio.ScanLines / ScanBytes
			scanner.Split(bufio.ScanLines)

			if scanner.Scan() { // true
				So(scanner.Text(), ShouldEqual, "foo")
			}
			if scanner.Scan() { // true
				So(scanner.Text(), ShouldEqual, "bar")
			}
			if scanner.Scan() { // true
				So(scanner.Text(), ShouldEqual, "baz")
			}

			So(scanner.Scan(), ShouldBeFalse)
		})
	})
}

func Test_Zip(t *testing.T) {

	zipPath := tmp_file + ".zip"
	file1 := zipPath + ".1"
	content1 := "string contents of file"
	file2 := zipPath + ".2"
	content2 := "\x61\x62\x63\n"

	Convey("", t, func() {
		f, err := os.Create(zipPath)
		defer os.Remove(zipPath)
		defer os.Remove(file1)
		defer os.Remove(file2)
		So(err, ShouldBeNil)
		defer f.Close()

		w := zip.NewWriter(f)

		var files = []struct {
			Name string
			Body string
		}{
			{file1, content1},
			{file2, content2},
		}

		Convey("zip file1, file2 -> tmp.txt.zip", func() {
			for _, file := range files {
				fwr, err := w.Create(file.Name)
				So(err, ShouldBeNil)
				_, err = fwr.Write([]byte(file.Body))
				So(err, ShouldBeNil)
			}
			w.Flush()
			w.Close() // 这里close() 下面才能openReader()

			Convey("unzip tmp.txt.zip -> file1, file2", func() {
				r, err := zip.OpenReader(zipPath)
				So(err, ShouldBeNil)
				defer r.Close()

				// 遍历解压每一个文件/文件夹
				for _, f := range r.Reader.File {
					zipFile, err := f.Open()
					So(err, ShouldBeNil)
					defer zipFile.Close()

					if f.FileInfo().IsDir() {
						os.MkdirAll(f.Name, f.Mode())
					} else {
						outFile, err := os.OpenFile(f.Name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
						So(err, ShouldBeNil)
						defer outFile.Close()

						_, err = io.Copy(outFile, zipFile)
						So(err, ShouldBeNil)
					}
				}

				Convey("check by sha256", func() {
					bs1, err := ioutil.ReadFile(file1)
					So(err, ShouldBeNil)
					So(len(bs1), ShouldEqual, len(content1))
					bs2, err := ioutil.ReadFile(file2)
					So(err, ShouldBeNil)
					So(len(bs2), ShouldEqual, len(content2))
					So(sha256.Sum256(bs2), ShouldEqual, sha256.Sum256([]byte(content2)))
				})
			})
		})
	})
}

func Test_Gz(t *testing.T) {
	gzPath := tmp_file + ".gz"
	content := "haha!\n"
	Convey("file -> gz", t, func() {
		f, err := os.Create(gzPath)
		So(err, ShouldBeNil)
		defer os.Remove(gzPath)

		w := gzip.NewWriter(f)
		w.Write([]byte(content))
		w.Close()

		Convey("gz -> file", func() {
			gzF, err := os.Open(gzPath)
			So(err, ShouldBeNil)
			defer os.Remove(tmp_file)

			r, err := gzip.NewReader(gzF)
			So(err, ShouldBeNil)
			defer r.Close()

			out, err := os.Create(tmp_file)
			So(err, ShouldBeNil)
			defer out.Close()

			_, err = io.Copy(out, r)
			So(err, ShouldBeNil)

			Convey("check by md5", func() {
				bs, err := ioutil.ReadFile(tmp_file)
				So(err, ShouldBeNil)
				So(md5.Sum(bs), ShouldEqual, md5.Sum([]byte(content)))

			})
		})
	})
}

func Test_TmpDir(t *testing.T) {
	Convey("", t, func() {
		dir, err := ioutil.TempDir("", "testtt")
		So(err, ShouldBeNil)
		f, err := ioutil.TempFile(dir, "file123")
		So(err, ShouldBeNil)

		// do something...

		defer f.Close()

		defer os.Remove(dir)
		defer os.Remove(f.Name())

		stat, err := os.Stat(dir)
		So(err, ShouldBeNil)
		So(stat.IsDir(), ShouldBeTrue)
	})
}
