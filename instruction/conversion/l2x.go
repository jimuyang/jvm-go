package conversion

import (
	"github.com/jimuyang/jvm-go/instruction"
	"github.com/jimuyang/jvm-go/rtda"
)

type L2I struct {
	instruction.NoOperandsInstruction
}

func (sf *L2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	stack.PushInt(int32(l))
}

type L2F struct {
	instruction.NoOperandsInstruction
}

func (sf *L2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	stack.PushFloat(float32(l))
}

type L2D struct {
	instruction.NoOperandsInstruction
}

func (sf *L2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	stack.PushDouble(float64(l))
}
