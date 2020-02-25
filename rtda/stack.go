package rtda

// Stack jvm栈
type Stack struct {
	maxSize uint   // 最多多少帧
	size    uint   // 栈当前大小
	_top    *Frame // 指向栈顶
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

func (sf *Stack) push(frame *Frame) {
	if sf.size >= sf.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if sf._top != nil {
		frame.lower = sf._top
	}
	sf._top = frame
	sf.size++
}

func (sf *Stack) pop() *Frame {
	if sf._top == nil {
		panic("Stack is Empty!")
	}
	top := sf._top
	sf._top = top.lower
	top.lower = nil
	sf.size--
	return top
}

func (sf *Stack) top() *Frame {
	if sf._top == nil {
		panic("Stack is Empty!")
	}
	return sf._top
}
