package math

import (
	"github.com/jimuyang/jvm-go/instruction"
	"github.com/jimuyang/jvm-go/rtda"
)

// int 左移
type ISHL struct {
	instruction.NoOperandsInstruction
}

func (sf *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	stack.PushInt(v1 << s)
}

// int 算术右移
type ISHR struct {
	instruction.NoOperandsInstruction
}

// int 逻辑右移
type IUSHR struct {
	instruction.NoOperandsInstruction
}

// long 左移
type LSHL struct {
	instruction.NoOperandsInstruction
}

func (sf *LSHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	stack.PushLong(v1 << s)
}

// long 算术右移
type LSHR struct {
	instruction.NoOperandsInstruction
}

// long 逻辑右移
type LUSHR struct {
	instruction.NoOperandsInstruction
}
