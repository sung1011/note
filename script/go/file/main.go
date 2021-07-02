package main

import (
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
	return ioutil.WriteFile(tf.path, bs, 0666)
}
func (tf *TheFile) ReadSimple() ([]byte, error) {
	return ioutil.ReadFile(tf.path)
}
