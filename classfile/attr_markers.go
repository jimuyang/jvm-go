package classfile

type DeprecatedAttribute struct {
	MarkerAttribute
}

type SyntheticAttribute struct {
	MarkerAttribute
}

// 仅起标记作用
type MarkerAttribute struct {
}

func (sf *MarkerAttribute) readInfo(reader *ClassReader) {
}
