package classfile

type UnparsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

func (sf *UnparsedAttribute) readInfo(reader *ClassReader) {
	sf.info = reader.readBytes(sf.length)
}
