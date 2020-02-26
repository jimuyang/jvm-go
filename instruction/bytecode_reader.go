package instruction

// BytecodeReader 字节码Reader
type BytecodeReader struct {
	code []byte
	pc   int
}

func (sf *BytecodeReader) Reset(code []byte, pc int) {
	sf.code = code
	sf.pc = pc
}

func (sf *BytecodeReader) ReadUint8() uint8 {
	i := sf.code[sf.pc]
	sf.pc++
	return i
}

func (sf *BytecodeReader) ReadByte() byte {
	return sf.ReadUint8()
}

func (sf *BytecodeReader) ReadInt8() int8 {
	return int8(sf.ReadUint8())
}

func (sf *BytecodeReader) ReadUint16() uint16 {
	byte1, byte2 := sf.ReadByte(), sf.ReadByte()
	return uint16(byte1)<<8 | uint16(byte2)
}

func (sf *BytecodeReader) ReadInt16() int16 {
	return int16(sf.ReadUint16())
}

func (sf *BytecodeReader) ReadInt32() int32 {
	byte1 := sf.ReadByte()
	byte2 := sf.ReadByte()
	byte3 := sf.ReadByte()
	byte4 := sf.ReadByte()
	return int32(byte1)<<24 | int32(byte2)<<16 | int32(byte3)<<8 | int32(byte4)
}
