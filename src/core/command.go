package core

import (
	"errors"
	"strconv"
	"strings"
)

type CommandType int

const (
	C_ARITHMETIC CommandType = iota
	C_PUSH
	C_POP
)

type Segment string

const (
	SEGMENT_LOCAL    Segment = "local"
	SEGMENT_ARGUMENT Segment = "argument"
	SEGMENT_THIS     Segment = "this"
	SEGMENT_THAT     Segment = "that"
	SEGMENT_CONSTANT Segment = "constant"
	SEGMENT_STATIC   Segment = "static"
	SEGMENT_POINTER  Segment = "pointer"
	SEGMENT_TEMP     Segment = "temp"
)

type Command struct {
	Instruction string
	Arg1        Segment
	Type        CommandType
	Arg2        int
	LineNumber  uint32
}

func newCommand(lineNumber uint32, line string) Command {
	commandType := getCommandType(&line)
	switch commandType {
	case C_ARITHMETIC:
		return newAritmeticCommand(lineNumber, line, commandType)
	case C_PUSH:
		return newPushCommand(lineNumber, line, commandType)
	case C_POP:
		return newPopCommand(lineNumber, line, commandType)
	}
	panic(errors.New("invalid command"))
}

func getCommandType(line *string) CommandType {
	switch strings.Split(*line, " ")[0] {
	case "push":
		return C_PUSH
	case "pop":
		return C_POP
	case "add", "sub", "neg", "eq", "gt", "lt", "and", "or", "not":
		return C_ARITHMETIC
	default:
		panic(errors.New("invalid command"))
	}
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
