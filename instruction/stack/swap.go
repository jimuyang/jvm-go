package stack

import (
	"github.com/jimuyang/jvm-go/instruction"
	"github.com/jimuyang/jvm-go/rtda"
)

type SWAP struct {
	instruction.NoOperandsInstruction
}

func (sf *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
