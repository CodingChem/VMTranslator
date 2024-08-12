package main

import (
	"os"
	"strings"

	"github.com/codingchem/VMTranslator/core"
)

func main() {
	args := os.Args
	inputFileName := args[1]
	baseFilaName := strings.Split(inputFileName, ".")[0]
	outputFileName := baseFilaName + ".asm"

	cmds := core.ParseFile(inputFileName)

	core.WriteCode(outputFileName, cmds)
}
