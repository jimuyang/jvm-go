package classfile

import "fmt"

// ClassFile class文件
type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	// constantPool ConstantPool
	// fields []*MemberInfo
	// methods []*MemberInfo
	// attributes []AttributeInfo
}

// Parse 将字节码解析为ClassFile
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			ok := false
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}
