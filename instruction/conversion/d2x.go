package conversion

import (
	"github.com/jimuyang/jvm-go/instruction"
	"github.com/jimuyang/jvm-go/rtda"
)

type D2F struct {
	instruction.NoOperandsInstruction
}

func (sf *D2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	stack.PushFloat(float32(d))
}

type D2I struct {
	instruction.NoOperandsInstruction
}

func (sf *D2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	stack.PushInt(int32(d))
}

type D2L struct {
	instruction.NoOperandsInstruction
}

func (sf *D2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	stack.PushLong(int64(d))
}
