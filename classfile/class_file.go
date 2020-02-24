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
	constantPool ConstantPool
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
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

// Print 打印ClassFile的重要信息
func (sf *ClassFile) Print() {
	fmt.Printf("version: %v.%v\n", sf.majorVersion, sf.minorVersion)
	fmt.Printf("constant counts: %v\n", len(sf.constantPool))
	fmt.Printf("access flags: 0x%x\n", sf.accessFlags)
	fmt.Printf("this class: %v\n", sf.Classname())
	fmt.Printf("super class: %v\n", sf.SuperClassName())
	fmt.Printf("interfaces: %v\n", sf.InterfacesNames())
	fmt.Printf("fields count: %v\n", len(sf.fields))
	for _, f := range sf.fields {
		fmt.Printf("   %s\n", f.Name())
	}

	fmt.Printf("methods count: %v\n", len(sf.methods))
	for _, m := range sf.methods {
		fmt.Printf("   %s\n", m.Name())
	}
}

func (sf *ClassFile) read(reader *ClassReader) {
	sf.readAndCheckMagic(reader)
	sf.readAndCheckVersion(reader)
	sf.constantPool = readConstantPool(reader)
	sf.accessFlags = reader.readUint16()
	sf.thisClass = reader.readUint16()
	sf.superClass = reader.readUint16()
	sf.interfaces = reader.readUint16s()
	sf.fields = readMembers(reader, sf.constantPool)
	sf.methods = readMembers(reader, sf.constantPool)
	sf.attributes = readAttributes(reader, sf.constantPool)
}

// 检查class文件的开头magic
func (sf *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

// class文件的版本 这里只支持java8：45.0-52.0
func (sf *ClassFile) readAndCheckVersion(reader *ClassReader) {
	sf.minorVersion = reader.readUint16()
	sf.majorVersion = reader.readUint16()
	switch sf.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if sf.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

func (sf *ClassFile) Classname() string {
	return sf.constantPool.getClassName(sf.thisClass)
}

func (sf *ClassFile) SuperClassName() string {
	if sf.superClass > 0 {
		return sf.constantPool.getClassName(sf.superClass)
	}
	return ""
}

func (sf *ClassFile) InterfacesNames() []string {
	interfacesNames := make([]string, len(sf.interfaces))
	for i, cpIndex := range sf.interfaces {
		interfacesNames[i] = sf.constantPool.getClassName(cpIndex)
	}
	return interfacesNames
}

// 常量池

// 类访问标志 类/接口 public/private

// getter
func (sf *ClassFile) MinorVersion() uint16 {
	return sf.minorVersion
}

// getter
func (sf *ClassFile) MajorVersion() uint16 {
	return sf.majorVersion
}
