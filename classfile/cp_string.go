package classfile

type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (sf *ConstantStringInfo) readInfo(reader *ClassReader) {
	sf.stringIndex = reader.readUint16()
}

func (sf *ConstantStringInfo) String() string {
	return sf.cp.getUtf8(sf.stringIndex)
}
