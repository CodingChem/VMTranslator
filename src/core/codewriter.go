package core

import "github.com/codingchem/VMTranslator/src/fileutils"

func WriteCode(fileName string, commands chan Command) {
	lines := make(chan string)
	go func() {
		defer close(lines)
		for command := range commands {
			for _, line := range translateCommand(command, fileName) {
				lines <- line
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
