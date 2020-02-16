package classfile

// MemberInfo 类的字段和方法
type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	// attributes []AttributeInfo
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		// attributes:
	}
}

func (sf *MemberInfo) Name() string {
	return sf.cp.getUtf8(sf.nameIndex)
}
func (sf *MemberInfo) Descriptor() string {
	return sf.cp.getUtf8(sf.descriptorIndex)
}
