package main

import (
	"errors"
	"fmt"
	"path"
)

// Visitor 访问者
type Visitor interface {
	Visit(IResourceFile) (string, error)
}

type IResourceFile interface {
	Accept(Visitor) (string, error)
}

func NewResourceFile(filepath string) (IResourceFile, error) {
	switch path.Ext(filepath) {
	case ".ppt":
		return &PPTFile{path: filepath}, nil
	case ".pdf":
		return &PdfFile{path: filepath}, nil
	default:
		return nil, fmt.Errorf("not found file type: %s", filepath)
	}
}

type PdfFile struct {
	path string
}

func (f *PdfFile) Accept(visitor Visitor) (string, error) {
	return visitor.Visit(f)
}

type PPTFile struct {
	path string
}

func (f *PPTFile) Accept(visitor Visitor) (string, error) {
	return visitor.Visit(f)
}

// Compressor 实现压缩功能
type Compressor struct{}

// Visit 实现访问者模式方法
// 我们可以发现由于没有函数重载, 我们只能通过断言来根据不同的类型调用不同函数
// 但是我们即使不采用访问者模式, 我们其实也是可以这么操作的
// 并且由于采用了类型断言, 所以如果需要操作的对象比较多的话, 这个函数其实也会膨胀的比较厉害
// 后续可以考虑按照命名约定使用 generate 自动生成代码
// 或者是使用反射简化
func (c *Compressor) Visit(r IResourceFile) (string, error) {
	switch f := r.(type) {
	case *PPTFile:
		return c.VisitPPTFile(f), nil
	case *PdfFile:
		return c.VisitPDFFile(f), nil
	default:
		return "", errors.New("not found file")
	}
}

func (c *Compressor) VisitPPTFile(f *PPTFile) string {
	return "ppt file"
}

func (c *Compressor) VisitPDFFile(f *PdfFile) string {
	return "pdf file"
}
