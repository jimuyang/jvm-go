package classpath

import (
	"io/ioutil"
	"path/filepath"
)

// DirEntry 目录入口
type DirEntry struct {
	absPath string
}

func newDirEntry(path string) *DirEntry {
	// 绝对路径
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absPath}
}

func (de *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(de.absPath, className)
	data, err := ioutil.ReadFile(fileName)
	return data, de, err
}

func (de *DirEntry) String() string {
	return de.absPath
}
