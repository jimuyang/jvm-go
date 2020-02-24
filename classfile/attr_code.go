package classfile

// 方法的字节码
type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (sf *CodeAttribute) readInfo(reader *ClassReader) {
	sf.maxStack = reader.readUint16()
	sf.maxLocals = reader.readUint16()

	codeLen := reader.readUint32()
	sf.code = reader.readBytes(codeLen)
	sf.exceptionTable = readExceptionTable(reader)
	sf.attributes = readAttributes(reader, sf.cp)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLen := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLen)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}
