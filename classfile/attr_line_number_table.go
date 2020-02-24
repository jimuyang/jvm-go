package classfile

type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (sf *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLen := reader.readUint16()
	sf.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLen)
	for i := range sf.lineNumberTable {
		sf.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}
