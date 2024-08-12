package core

import (
	"fmt"
	"strings"
)

func newAritmeticCommand(lineNumber uint32, line string, commandType CommandType) Command {
	return Command{
		Instruction: strings.Split(line, " ")[0],
		Type:        commandType,
		LineNumber:  lineNumber,
	}
}

func translateArithmetic(command Command) []string {
	switch command.Instruction {
	case "add":
		return add()
	case "sub":
		return sub()
	case "neg":
		return neg()
	case "eq":
		return eq(command.LineNumber)
	case "gt":
		return gt(command.LineNumber)
	case "lt":
		return lt(command.LineNumber)
	case "and":
		return and()
	case "or":
		return or()
	case "not":
		return not()
	default:
		panic("Invalid arithmetic instruction")
	}
}

func add() []string {
	return []string{
		"//add",
		"@SP",
		"AM=M-1",
		"D=M",
		"A=A-1",
		"M=D+M",
	}
}

func sub() []string {
	return []string{
		"//sub",
		"@SP",
		"AM=M-1",
		"D=M",
		"A=A-1",
		"M=M-D",
	}
}

func neg() []string {
	return []string{
		"//neg",
		"@SP",
		"A=M-1",
		"M=-M",
	}
}

func eq(lineNumber uint32) []string {
	return []string{
		"//eq",
		"@SP",
		"AM=M-1",
		"D=M",
		"A=A-1",
		"D=M-D",
		fmt.Sprintf("@EQ_TRUE_%d", lineNumber),
		"D;JEQ",
		"D=0",
		fmt.Sprintf("@EQ_END_%d", lineNumber),
		"0;JMP",
		fmt.Sprintf("(EQ_TRUE_%d)", lineNumber),
		"D=-1",
		fmt.Sprintf("(EQ_END_%d)", lineNumber),
		"@SP",
		"A=M-1",
		"M=D",
	}
}

func gt(lineNumber uint32) []string {
	return []string{
		"//gt",
		"@SP",
		"AM=M-1",
		"D=M",
		"A=A-1",
		"D=M-D",
		fmt.Sprintf("@GT_TRUE_%d", lineNumber),
		"D;JGT",
		"D=0",
		fmt.Sprintf("@GT_END_%d", lineNumber),
		"0;JMP",
		fmt.Sprintf("(GT_TRUE_%d)", lineNumber),
		"D=-1",
		fmt.Sprintf("(GT_END_%d)", lineNumber),
		"@SP",
		"A=M-1",
		"M=D",
	}
}

func lt(lineNumber uint32) []string {
	return []string{
		"//lt",
		"@SP",
		"AM=M-1",
		"D=M",
		"A=A-1",
		"D=M-D",
		fmt.Sprintf("@LT_TRUE_%d", lineNumber),
		"D;JLT",
		"D=0",
		fmt.Sprintf("@LT_END_%d", lineNumber),
		"0;JMP",
		fmt.Sprintf("(LT_TRUE_%d)", lineNumber),
		"D=-1",
		fmt.Sprintf("(LT_END_%d)", lineNumber),
		"@SP",
		"A=M-1",
		"M=D",
	}
}

func and() []string {
	return []string{
		"//and",
		"@SP",
		"AM=M-1",
		"D=M",
		"A=A-1",
		"M=D&M",
	}
}

func or() []string {
	return []string{
		"//or",
		"@SP",
		"AM=M-1",
		"D=M",
		"A=A-1",
		"M=D|M",
	}
}

func not() []string {
	return []string{
		"//not",
		"@SP",
		"A=M-1",
		"M=!M",
	}
}
