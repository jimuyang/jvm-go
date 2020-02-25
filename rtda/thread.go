package rtda

// Thread jvm线程
type Thread struct {
	pc    int    // program counter
	stack *Stack // 线程帧栈
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

func (sf *Thread) PC() int {
	return sf.pc
}

func (sf *Thread) SetPC(pc int) {
	sf.pc = pc
}

func (sf *Thread) PushFrame(frame *Frame) {
	sf.stack.push(frame)
}

func (sf *Thread) PopFrame() *Frame {
	return sf.stack.pop()
}

func (sf *Thread) CurrentFrame() *Frame {
	return sf.stack.top()
}
