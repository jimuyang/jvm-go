package classfile

type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (sf *ConstantClassInfo) readInfo(reader *ClassReader) {
	sf.nameIndex = reader.readUint16()
}

func (sf *ConstantClassInfo) Name() string {
	return sf.cp.getUtf8(sf.nameIndex)
}
