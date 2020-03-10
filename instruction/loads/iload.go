package loads

import (
	"github.com/jimuyang/jvm-go/instruction"
	"github.com/jimuyang/jvm-go/rtda"
)

type ILOAD struct {
	instruction.Index8Instruction
}

func (sf *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, sf.Index)
}

func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

type ILOAD_0 struct {
	instruction.NoOperandsInstruction
}

func (sf *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}

type ILOAD_1 struct {
	instruction.NoOperandsInstruction
}

func (sf *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

type ILOAD_2 struct {
	instruction.NoOperandsInstruction
}

func (sf *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

type ILOAD_3 struct {
	instruction.NoOperandsInstruction
}

func (sf *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}
