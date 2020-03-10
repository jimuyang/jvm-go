// Package constants 这一系列指令将隐含在操作码中的常量值推入操作数栈顶
package constants

import (
	"github.com/jimuyang/jvm-go/instruction"
	"github.com/jimuyang/jvm-go/rtda"
)

type ACONST_NULL struct {
	instruction.NoOperandsInstruction
}

type DCONST_0 struct {
	instruction.NoOperandsInstruction
}
type DCONST_1 struct {
	instruction.NoOperandsInstruction
}

type FCONST_0 struct {
	instruction.NoOperandsInstruction
}
type FCONST_1 struct {
	instruction.NoOperandsInstruction
}
type FCONST_2 struct {
	instruction.NoOperandsInstruction
}

type ICONST_0 struct {
	instruction.NoOperandsInstruction
}
type ICONST_1 struct {
	instruction.NoOperandsInstruction
}
type ICONST_2 struct {
	instruction.NoOperandsInstruction
}
type ICONST_3 struct {
	instruction.NoOperandsInstruction
}
type ICONST_4 struct {
	instruction.NoOperandsInstruction
}
type ICONST_5 struct {
	instruction.NoOperandsInstruction
}
type ICONST_M1 struct {
	instruction.NoOperandsInstruction
}

type LCONST_0 struct {
	instruction.NoOperandsInstruction
}
type LCONST_1 struct {
	instruction.NoOperandsInstruction
}

func (sf *ACONST_NULL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
}

func (sf *DCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

func (sf *DCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(1.0)
}

func (sf *FCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(0.0)
}

func (sf *FCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(1.0)
}

func (sf *FCONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(2.0)
}

func (sf *ICONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(0)
}

func (sf *ICONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(1)
}

func (sf *ICONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(2)
}

func (sf *ICONST_3) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(3)
}

func (sf *ICONST_4) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(4)
}

func (sf *ICONST_5) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(5)
}

func (sf *ICONST_M1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(-1)
}

func (sf *LCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(0)
}

func (sf *LCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(1)
}
