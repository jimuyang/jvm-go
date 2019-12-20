package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

// Entry 入口
type Entry interface {
	// 返回读取到class字节码
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		// 多个路径
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, "*") {
		// 有通配符
		return newWildcardEntry(path)
	}

	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		// 是一个jar包
		return newJarEntry(path)
	}

	// 认为是一个目录
	return newDirEntry(path)
}
