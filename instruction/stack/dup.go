package stack

import (
	"jvm-go/instruction"
	"jvm-go/rtda"
)

type DUP struct {
	instruction.NoOperandsInstruction
}

type DUP_X1 struct {
	instruction.NoOperandsInstruction
}

type DUP_X2 struct {
	instruction.NoOperandsInstruction
}

type DUP2 struct {
	instruction.NoOperandsInstruction
}

type DUP2_X1 struct {
	instruction.NoOperandsInstruction
}

type DUP2_X2 struct {
	instruction.NoOperandsInstruction
}

func (sf *DUP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}
