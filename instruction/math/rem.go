package math

import (
	"github.com/jimuyang/jvm-go/instruction"
	"github.com/jimuyang/jvm-go/rtda"
	"math"
)

type DREM struct {
	instruction.NoOperandsInstruction
}

func (sf *DREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := math.Mod(v1, v2)
	stack.PushDouble(result)
}

type FREM struct {
	instruction.NoOperandsInstruction
}

func (sf *FREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := math.Mod(float64(v1), float64(v2))
	stack.PushFloat(float32(result))
}

type IREM struct {
	instruction.NoOperandsInstruction
}

func (sf *IREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 % v2
	stack.PushInt(result)
}

type LREM struct {
	instruction.NoOperandsInstruction
}

func (sf *LREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 % v2
	stack.PushLong(result)
}
