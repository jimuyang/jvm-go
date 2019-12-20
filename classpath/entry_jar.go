package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

// JarEntry Jar包
type JarEntry struct {
	absPath string
}

func newJarEntry(path string) *JarEntry {
	// 绝对路径
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &JarEntry{absPath}
}

func (je *JarEntry) readClass(className string) ([]byte, Entry, error) {
	// 打开zip包
	reader, err := zip.OpenReader(je.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer reader.Close()

	for _, f := range reader.File {
		if f.Name == className {
			// 找到文件
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()

			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, je, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (je *JarEntry) String() string {
	return je.absPath
}
