package rtda

import (
	"math"
)

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
	sf.slots[sf.size].num = val
	sf.size++
}

func (sf *OperandStack) PopInt() int32 {
	sf.size--
	return sf.slots[sf.size].num
}

func (sf *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	sf.PushInt(int32(bits))
}

func (sf *OperandStack) PopFloat() float32 {
	bits := uint32(sf.PopInt())
	return math.Float32frombits(bits)
}

func (sf *OperandStack) PushLong(val int64) {
	sf.slots[sf.size].num = int32(val)
	sf.slots[sf.size+1].num = int32(val >> 32)
	sf.size += 2
}

func (sf *OperandStack) PopLong() int64 {
	sf.size -= 2
	low := uint32(sf.slots[sf.size].num)
	high := uint32(sf.slots[sf.size+1].num)
	return int64(high)<<32 | int64(low)
}

func (sf *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	sf.PushLong(int64(bits))
}

func (sf *OperandStack) PopDouble() float64 {
	bits := uint64(sf.PopLong())
	return math.Float64frombits(bits)
}

func (sf *OperandStack) PushRef(val *Object) {
	sf.slots[sf.size].ref = val
	sf.size++
}

func (sf *OperandStack) PopRef() *Object {
	sf.size--
	ref := sf.slots[sf.size].ref
	sf.slots[sf.size].ref = nil // 垃圾回收
	return ref
}
