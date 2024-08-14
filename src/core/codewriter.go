package core

import "github.com/codingchem/VMTranslator/src/fileutils"

type VMFile struct {
	Commands chan Command
	Name     string
}

func WriteCode(VMFiles []VMFile, fileName string) {
	lines := make(chan string)
	go func() {
		defer close(lines)
		for _, file := range VMFiles {
			fileName := file.Name
			for command := range file.Commands {
				for _, line := range translateCommand(command, fileName) {
					lines <- line
				}
			}
		}
	}()
	fileutils.WriteFileLines(fileName, lines)
}

func translateCommand(command Command, fileName string) []string {
	switch command.Type {
	case C_ARITHMETIC:
		return translateArithmetic(command)
	case C_PUSH:
		return translatePush(command, fileName)
	case C_POP:
		return translatePop(command, fileName)
	default:
		panic("invalid command")
	}
}
