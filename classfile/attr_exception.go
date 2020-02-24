package classfile

// 方法抛出的异常表
type ExceptionAttribute struct {
	exceptionIndexTable []uint16
}

func (sf *ExceptionAttribute) readInfo(reader *ClassReader) {
	sf.exceptionIndexTable = reader.readUint16s()
}
func (sf *ExceptionAttribute) ExceptionIndexTable() []uint16 {
	return sf.exceptionIndexTable
}
