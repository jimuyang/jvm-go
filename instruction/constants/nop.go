package constants

import (
	"github.com/jimuyang/jvm-go/instruction"
	"github.com/jimuyang/jvm-go/rtda"
)

// NOP 该指令什么也不做
type NOP struct {
	instruction.NoOperandsInstruction
}

func (sf *NOP) Execute(frame *rtda.Frame) {
}
