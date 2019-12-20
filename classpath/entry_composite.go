package classpath

import (
	"errors"
	"strings"
)

// CompositeEntry 组合
type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	ce := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		ce = append(ce, entry)
	}
	return ce
}

func (ce CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range ce {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, err
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (ce CompositeEntry) String() string {
	strs := make([]string, len(ce))
	for i, entry := range ce {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}
