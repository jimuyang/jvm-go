package rtda

// Frame jvm栈帧
type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
}

func newFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (sf *Frame) OperandStack() *OperandStack {
	return sf.operandStack
}

func (sf *Frame) LocalVars() LocalVars {
	return sf.localVars
}
