package constants

import (
	"jvm-go/instruction"
	"jvm-go/rtda"
)

// NOP 该指令什么也不做
type NOP struct {
	instruction.NoOperandsInstruction
}

func (sf *NOP) Execute(frame *rtda.Frame) {
}
