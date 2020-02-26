package instruction

import (
	"jvm-go/rtda"
)

// Instruction 指令接口
type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

// NoOperandsInstruction 无操作数的指令
type NoOperandsInstruction struct{}

func (sf *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {}

// BranchInstruction 跳转指令
type BranchInstruction struct {
	Offset int // 跳转偏移量
}

func (sf *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	sf.Offset = int(reader.ReadInt16())
}

// Index8Instruction 按照索引存取局部变量表
type Index8Instruction struct {
	Index uint
}

func (sf *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	sf.Index = uint(reader.ReadUint8())
}

// Index16Instruction 操作常量池索引（2字节操作数）
type Index16Instruction struct {
	Index uint
}

func (sf *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	sf.Index = uint(reader.ReadUint16())
}
