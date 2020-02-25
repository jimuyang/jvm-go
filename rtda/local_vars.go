package rtda

import (
	"math"
)

// LocalVars 局部变量表
type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

func (sf LocalVars) GetInt(index uint) int32 {
	return sf[index].num
}

func (sf LocalVars) SetInt(index uint, val int32) {
	sf[index].num = val
}

func (sf LocalVars) GetFloat(index uint) float32 {
	bits := uint32(sf[index].num)
	return math.Float32frombits(bits)
}

func (sf LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	sf[index].num = int32(bits)
}

func (sf LocalVars) GetLong(index uint) int64 {
	low := uint32(sf[index].num)
	high := uint32(sf[index+1].num)
	return int64(high<<32) | int64(low)
}

func (sf LocalVars) SetLong(index uint, val int64) {
	sf[index].num = int32(val)
	sf[index+1].num = int32(val >> 32)
}

func (sf LocalVars) GetDouble(index uint) float64 {
	bits := uint64(sf.GetLong(index))
	return math.Float64frombits(bits)
}

func (sf LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	sf.SetLong(index, int64(bits))
}

func (sf LocalVars) GetRef(index uint) *Object {
	return sf[index].ref
}

func (sf LocalVars) SetRef(index uint, ref *Object) {
	sf[index].ref = ref
}
