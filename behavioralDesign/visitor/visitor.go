/*
 * @Author: bsun
 * @Date: 2024-07-02 13:13:10
 * @Last Modified by: bsun
 * @Last Modified time: 2024-07-02 16:13:57
 */
package visitor

import (
	"fmt"
	"path"
)

// 访问者 接口集合
type Vistor interface {
	Visit(IResourceFile) error
}

// 资源文件
type IResourceFile interface {
	Accept(Vistor) error
}

// 生成资源器
func NewResourceFile(filePath string) (IResourceFile, error) {
	switch path.Ext(filePath) {
	case ".ppt":
		return &PPTfFile{path: filePath}, nil
	case ".pdf":
		return &PdFile{path: filePath}, nil
	}
	return nil, fmt.Errorf("not found file type: %s", filePath)
}

// 资源文件
type PdFile struct {
	path string
}

func (f *PdFile) Accept(vst Vistor) error {
	return vst.Visit(f)
}

type PPTfFile struct {
	path string
}

func (f *PPTfFile) Accept(vst Vistor) error {
	return vst.Visit(f)
}

// compressor实现压缩功能
type Compressor struct{}

func (c *Compressor) Visit(r IResourceFile) error {
	switch f := r.(type) {
	case *PPTfFile:
		return c.VisitPPTFile(f)
	case *PdFile:
		return c.VisitPDFile(f)
	default:
		return fmt.Errorf("not found resource type:%#v", r)
	}
}

func (c *Compressor) VisitPPTFile(f *PPTfFile) error {
	fmt.Println("this is ppt file")
	return nil
}

func (c *Compressor) VisitPDFile(f *PdFile) error {
	fmt.Println("this is pdf file")
	return nil
}
