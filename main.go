package main

import (
	"fmt"
	"jvm-go/classfile"
	"jvm-go/classpath"
	"strings"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 1.0.0")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	fmt.Printf("classpath:%s class:%s args:%v \n", cmd.cpOption, cmd.class, cmd.args)
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)

	className := strings.Replace(cmd.class, ".", "/", -1)
	classFile := loadClass(className, cp)
	classFile.Print()
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("cannot find or load main class %s \n", className)
		panic(err)
	}
	fmt.Printf("class data: %v \n", classData)

	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}
