package classfile

type ConstantMemberInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (sf *ConstantMemberInfo) readInfo(reader *ClassReader) {
	sf.classIndex = reader.readUint16()
	sf.nameAndTypeIndex = reader.readUint16()
}

func (sf *ConstantMemberInfo) Classname() string {
	return sf.cp.getClassName(sf.classIndex)
}

func (sf *ConstantMemberInfo) NameAndDescriptor() (string, string) {
	return sf.cp.getNameAndType(sf.nameAndTypeIndex)
}

type ConstantFieldRefInfo struct {
	ConstantMemberInfo
}
type ConstantMethodRefInfo struct {
	ConstantMemberInfo
}

type ConstantInterfaceMethodRefInfo struct {
	ConstantMemberInfo
}
