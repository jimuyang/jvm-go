package classfile

type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (sf *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	sf.nameIndex = reader.readUint16()
	sf.descriptorIndex = reader.readUint16()
}
