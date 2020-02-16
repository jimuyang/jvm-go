package classfile

import "math"

// ConstantIntegerInfo int short byte 使用4字节
type ConstantIntegerInfo struct {
	val int32
}

func (sf *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	sf.val = int32(bytes)
}

// ConstantFloatInfo 单精度数 使用4byte
type ConstantFloatInfo struct {
	val float32
}

func (sf *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	sf.val = math.Float32frombits(bytes)
}

// ConstantLongInfo long 使用8byte
type ConstantLongInfo struct {
	val int64
}

func (sf *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	sf.val = int64(bytes)
}

// ConstantDoubleInfo double 双精度数 使用8byte
type ConstantDoubleInfo struct {
	val float64
}

func (sf *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	sf.val = math.Float64frombits(bytes)
}
