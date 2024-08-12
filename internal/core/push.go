package core

import (
	"fmt"
	"strings"
)

func newPushCommand(lineNumber uint32, line string, commandType CommandType) Command {
	splitted := strings.Split(line, " ")
	return Command{
		Instruction: splitted[0],
		Arg1:        Segment(splitted[1]),
		Arg2:        toInt(splitted[2]),
		Type:        commandType,
		LineNumber:  lineNumber,
	}
}

func translatePush(command Command, fileName string) []string {
	switch command.Arg1 {
	case SEGMENT_CONSTANT:
		return pushConstant(command)
	case SEGMENT_LOCAL:
		return pushLocal(command)
	case SEGMENT_ARGUMENT:
		return pushArgument(command)
	case SEGMENT_THIS:
		return pushThis(command)
	case SEGMENT_THAT:
		return pushThat(command)
	case SEGMENT_TEMP:
		return pushTemp(command)
	case SEGMENT_POINTER:
		return pushPointer(command)
	case SEGMENT_STATIC:
		return pushStatic(command, fileName)
	default:
		panic("Invalid segment for push")
	}
}

func pushConstant(command Command) []string {
	return []string{
		fmt.Sprintf("//push constant %d", command.Arg2),
		fmt.Sprintf("@%d", command.Arg2),
		"D=A",
		"@SP",
		"A=M",
		"M=D",
		"@SP",
		"M=M+1",
	}
}

func pushLocal(command Command) []string {
	return []string{
		fmt.Sprintf("// push local %d", command.Arg2),
		"@LCL",
		"D=M",
		fmt.Sprintf("@%d", command.Arg2),
		"A=D+A",
		"D=M",
		"@SP",
		"A=M",
		"M=D",
	}
}

func pushArgument(command Command) []string {
	return []string{
		fmt.Sprintf("// push argument %d", command.Arg2),
		"@ARG",
		"D=M",
		fmt.Sprintf("@%d", command.Arg2),
		"A=D+A",
		"D=M",
		"@SP",
		"A=M",
		"M=D",
	}
}

func pushThis(command Command) []string {
	return []string{
		fmt.Sprintf("// push this %d", command.Arg2),
		fmt.Sprintf("@%d", command.Arg2),
		"D=A",
		"@THIS",
		"A=D+A",
		"D=M",
		"@SP",
		"A=M",
		"M=D",
		"@SP",
		"M=M+1",
	}
}

func pushThat(command Command) []string {
	return []string{
		fmt.Sprintf("// push that %d", command.Arg2),
		fmt.Sprintf("@%d", command.Arg2),
		"D=A",
		"@THAT",
		"A=D+A",
		"D=M",
		"@SP",
		"A=M",
		"M=D",
		"@SP",
		"M=M+1",
	}
}

func pushTemp(command Command) []string {
	// temp 0 is at R5
	address := command.Arg2 + 5
	if address > 12 || address < 5 {
		panic("Invalid temp address")
	}
	return []string{
		fmt.Sprintf("// push temp %d", command.Arg2),
		fmt.Sprintf("@%d", address),
		"D=M",
		"@SP",
		"A=M",
		"M=D",
		"@SP",
		"M=M+1",
	}
}

func pushPointer(command Command) []string {
	// pointer 0 is at THIS, pointer 1 is at THAT
	var segment string
	if command.Arg2 == 0 {
		segment = "@THIS"
	} else {
		segment = "@THAT"
	}
	return []string{
		fmt.Sprintf("// push pointer %d", command.Arg2),
		segment,
		"D=M",
		"@SP",
		"A=M",
		"M=D",
		"@SP",
		"M=M+1",
	}
}

func pushStatic(command Command, fileName string) []string {
	return []string{
		fmt.Sprintf("// push static %d", command.Arg2),
		fmt.Sprintf("@%s.%d", strings.Join(strings.Split(fileName, "/"), "."), command.Arg2),
		"D=M",
		"@SP",
		"M=M+1",
		"A=M-1",
		"M=D",
	}
}
