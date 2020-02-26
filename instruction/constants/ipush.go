package constants

import (
	"jvm-go/instruction"
	"jvm-go/rtda"
)

type BIPUSH struct {
	val int8
}

func (sf *BIPUSH) FetchOperands(reader *instruction.BytecodeReader) {
	sf.val = reader.ReadInt8()
}

func (sf *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(sf.val)
	frame.OperandStack().PushInt(i)
}

type SIPUSH struct {
	val int16
}

func (sf *SIPUSH) FetchOperands(reader *instruction.BytecodeReader) {
	sf.val = reader.ReadInt16()
}

func (sf *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(sf.val)
	frame.OperandStack().PushInt(i)
}
