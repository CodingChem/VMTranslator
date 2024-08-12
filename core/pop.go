package core

import (
	"fmt"
	"strings"
)

func newPopCommand(lineNumber uint32, line string, commandType CommandType) Command {
	splitted := strings.Split(line, " ")
	return Command{
		Instruction: splitted[0],
		Arg1:        Segment(splitted[1]),
		Arg2:        toInt(splitted[2]),
		Type:        commandType,
		LineNumber:  lineNumber,
	}
}

func translatePop(command Command, fileName string) []string {
	switch command.Arg1 {
	case SEGMENT_LOCAL:
		return popLocal(command)
	case SEGMENT_ARGUMENT:
		return popArgument(command)
	case SEGMENT_THIS:
		return popThis(command)
	case SEGMENT_THAT:
		return popThat(command)
	case SEGMENT_TEMP:
		return popTemp(command)
	case SEGMENT_POINTER:
		return popPointer(command)
	case SEGMENT_STATIC:
		return popStatic(command, fileName)
	case SEGMENT_CONSTANT:
		panic("Cannot pop to constant")
	default:
		panic("Invalid segment for pop")
	}
}

func popLocal(command Command) []string {
	return []string{
		fmt.Sprintf("// pop local %d", command.Arg2),
		"@SP",
		"AM=M-1",
		"D=M",
		"@LCL",
		"D=D+M",
		fmt.Sprintf("@%d", command.Arg2),
		"D=D+A",
		"@SP",
		"A=M",
		"A=M",
		"A=D-A",
		"M=D-A",
	}
}

func popArgument(command Command) []string {
	return []string{
		fmt.Sprintf("// pop argument %d", command.Arg2),
		"@SP",
		"AM=M-1",
		"D=M",
		"@ARG",
		"D=D+M",
		fmt.Sprintf("@%d", command.Arg2),
		"D=D+A",
		"@SP",
		"A=M",
		"A=M",
		"A=D-A",
		"M=D-A",
	}
}

func popThis(command Command) []string {
	return []string{
		fmt.Sprintf("// pop this %d", command.Arg2),
		"@SP",
		"AM=M-1",
		"D=M",
		"@THIS",
		"D=D+M",
		fmt.Sprintf("@%d", command.Arg2),
		"D=D+A",
		"@SP",
		"A=M",
		"A=D-M",
		"M=D-A",
	}
}

func popThat(command Command) []string {
	return []string{
		fmt.Sprintf("// pop that %d", command.Arg2),
		"@SP",
		"AM=M-1",
		"D=M",
		"@THAT",
		"D=D+M",
		fmt.Sprintf("@%d", command.Arg2),
		"D=D+A",
		"@SP",
		"A=M",
		"A=D-M",
		"M=D-A",
	}
}

func popTemp(command Command) []string {
	address := command.Arg2 + 5
	if address > 12 || address < 5 {
		panic("Invalid temp address")
	}
	return []string{
		fmt.Sprintf("// pop temp %d", command.Arg2),
		"@SP",
		"AM=M-1",
		"D=M",
		fmt.Sprintf("@%d", address),
		"M=D",
	}
}

func popPointer(command Command) []string {
	var segment string
	if command.Arg2 == 0 {
		segment = "@THIS"
	} else {
		segment = "@THAT"
	}
	return []string{
		fmt.Sprintf("// pop pointer %d", command.Arg2),
		"@SP",
		"AM=M-1",
		"D=M",
		segment,
		"M=D",
	}
}

func popStatic(command Command, fileName string) []string {
	return []string{
		fmt.Sprintf("// pop static %d", command.Arg2),
		"@SP",
		"AM=M-1",
		"D=M",
		fmt.Sprintf("@%s.%d", strings.Join(strings.Split(fileName, "/"), "."), command.Arg2),
		"M=D",
	}
}
