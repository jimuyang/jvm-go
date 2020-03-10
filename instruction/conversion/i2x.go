package conversion

import (
	"github.com/jimuyang/jvm-go/instruction"
	"github.com/jimuyang/jvm-go/rtda"
)

type I2L struct {
	instruction.NoOperandsInstruction
}

func (sf *I2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	stack.PushLong(int64(i))
}

type I2F struct {
	instruction.NoOperandsInstruction
}

func (sf *I2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	stack.PushFloat(float32(i))
}

type I2D struct {
	instruction.NoOperandsInstruction
}

func (sf *I2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	stack.PushDouble(float64(i))
}
