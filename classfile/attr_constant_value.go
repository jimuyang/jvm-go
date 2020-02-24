package classfile

// 用于表示常量表达式的值
type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (sf *ConstantValueAttribute) readInfo(reader *ClassReader) {
	sf.constantValueIndex = reader.readUint16()
}

func (sf *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return sf.constantValueIndex
}
