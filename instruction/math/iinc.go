package math

import (
	"github.com/jimuyang/jvm-go/instruction"
	"github.com/jimuyang/jvm-go/rtda"
)

// 给局部变量表中的int变量增加常量值
type IINC struct {
	Index uint
	Const int32
}

func (sf *IINC) FetchOperands(reader *instruction.BytecodeReader) {
	sf.Index = uint(reader.ReadUint8())
	sf.Const = int32(reader.ReadInt8())
}

func (sf *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(sf.Index)
	val += sf.Const
	localVars.SetInt(sf.Index, val)
}
