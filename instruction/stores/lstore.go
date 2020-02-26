package stores

import (
	"jvm-go/instruction"
	"jvm-go/rtda"
)

type LSTORE struct {
	instruction.Index8Instruction
}
type LSTORE_0 struct {
	instruction.NoOperandsInstruction
}
type LSTORE_1 struct {
	instruction.NoOperandsInstruction
}
type LSTORE_2 struct {
	instruction.NoOperandsInstruction
}
type LSTORE_3 struct {
	instruction.NoOperandsInstruction
}

func _lstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}

func (sf *LSTORE) Execute(frame *rtda.Frame) {
	_lstore(frame, sf.Index)
}
func (sf *LSTORE_0) Execute(frame *rtda.Frame) {
	_lstore(frame, 0)
}
func (sf *LSTORE_1) Execute(frame *rtda.Frame) {
	_lstore(frame, 1)
}
func (sf *LSTORE_2) Execute(frame *rtda.Frame) {
	_lstore(frame, 2)
}
func (sf *LSTORE_3) Execute(frame *rtda.Frame) {
	_lstore(frame, 3)
}
