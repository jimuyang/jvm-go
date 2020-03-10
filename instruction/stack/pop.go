package stack

import (
	"github.com/jimuyang/jvm-go/instruction"
	"github.com/jimuyang/jvm-go/rtda"
)

type POP struct {
	instruction.NoOperandsInstruction
}

func (sf *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

type POP2 struct {
	instruction.NoOperandsInstruction
}

func (sf *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
