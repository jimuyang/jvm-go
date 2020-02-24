package classfile

// AttributeInfo 属性表
type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

// 读取属性表
func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attrCount := reader.readUint16()
	attrs := make([]AttributeInfo, attrCount)
	for i := range attrs {
		attrs[i] = readAttribute(reader, cp)
	}
	return attrs
}

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attrNameIndex)
	attrLen := reader.readUint32()
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}

func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	// jvm预定义了23种属性
	switch attrName {
	case "Code":
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return &DeprecatedAttribute{}
	case "Exceptions":
		return &ExceptionAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	// case "LocalVariableTable":
	// return &LocalVariableTable
	case "SourceFile":
		return &SourceFileAttribute{}
	case "Synthetic":
		return &SyntheticAttribute{}
	default:
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
}
