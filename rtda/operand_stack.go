package rtda

// OperandStack 操作数栈
type OperandStack struct {
	size  uint // 记录栈顶
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

func (sf *OperandStack) PushInt(val int32) {
}

func (sf *OperandStack) PopInt() int32 {
}

func (sf *OperandStack) PushFloat(val float32) {
}

func (sf *OperandStack) PopFloat() float32 {
}

func (sf *OperandStack) PushLong(val int64) {
}

func (sf *OperandStack) PopLong() int64 {
}

func (sf *OperandStack) PushDouble(val float64) {
}

func (sf *OperandStack) PopDouble() float64 {
}

func (sf *OperandStack) PushRef(val *Object) {
}

func (sf *OperandStack) PopRef() *Object {
}
