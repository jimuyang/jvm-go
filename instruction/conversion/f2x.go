package conversion

import (
	"github.com/jimuyang/jvm-go/instruction"
	"github.com/jimuyang/jvm-go/rtda"
)

type F2I struct {
	instruction.NoOperandsInstruction
}

func (sf *F2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	stack.PushInt(int32(f))
}

type F2D struct {
	instruction.NoOperandsInstruction
}

func (sf *F2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	stack.PushDouble(float64(f))
}

type F2L struct {
	instruction.NoOperandsInstruction
}

func (sf *F2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	stack.PushLong(int64(f))
}
