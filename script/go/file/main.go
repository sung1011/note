package main

// https://colobu.com/2016/10/12/go-file-operations/

import (
	"io"
	"io/ioutil"
	"os"
)

var doc_file = "doc.txt"
var tmp_file = "tmp.txt"

type TheFile struct {
	path string
}

func New(path string) *TheFile {
	return &TheFile{path}
}

func (tf *TheFile) Create() (*os.File, error) {
	f, err := os.Create(tf.path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return f, nil
}

func (tf *TheFile) Exist() bool {
	_, err := os.Stat(tf.path)

	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func (tf *TheFile) Remove() error {
	return os.Remove(tf.path)
}

func (tf *TheFile) Truncate(n int64) error {
	return os.Truncate(tf.path, n)
}

func (tf *TheFile) WriteSimple(bs []byte) error {
	// 覆盖写 os.O_WRONLY|os.O_CREATE|os.O_TRUNC
	return ioutil.WriteFile(tf.path, bs, 0666)
}
func (tf *TheFile) ReadSimple() ([]byte, error) {
	return ioutil.ReadFile(tf.path)
}

func (tf *TheFile) Mv(new string) error {
	old := tf.path
	return os.Rename(old, new)
}

func (tf *TheFile) Open(flag int, mode os.FileMode) (*os.File, error) {
	// flag:
	//   os.O_RDONLY // 只读
	//   os.O_WRONLY // 只写
	//   os.O_RDWR // 读写
	//   os.O_APPEND // 往文件中添建（Append）
	//   os.O_CREATE // 如果文件不存在则先创建
	//   os.O_TRUNC // 文件打开时裁剪文件
	//   os.O_EXCL // 和O_CREATE一起使用，文件不能存在
	//   os.O_SYNC // 以同步I/O的方式打开

	// mode: 如0666
	//   d=isdir, u=own, g=group, o=other
	//   r = 4
	//   w = 2
	//   x = 1
	//   rwx = 4 + 2 + 1 = 7
	//   rw = 4 + 2 = 6
	//   rx = 4 +1 = 5

	// os.Open(name) == OpenFile(name, O_RDONLY, 0)  只读

	// os.Create(name) == OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)  读|写|空|创建

	f, err := os.OpenFile(tf.path, flag, mode)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return f, nil
}

func (tf *TheFile) Copy(target string) error {
	src, err := os.Open(tf.path)
	if err != nil {
		return err
	}
	tgt, err := os.Create(target)
	if err != nil {
		return err
	}
	// 覆盖式copy
	_, err = io.Copy(tgt, src)
	if err != nil {
		return err
	}
	return tgt.Sync()
}
