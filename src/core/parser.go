package core

import (
	"github.com/codingchem/VMTranslator/src/fileutils"
)

func ParseFile(fileName string) chan Command {
	out := make(chan Command)
	lines, err := fileutils.ReadFileLines(fileName)
	if err != nil {
		panic(err)
	}
	var lineNumber uint32 = 0
	go func() {
		defer close(out)
		for line := range lines {
			if !isCommentOrEmptyLine(&line) {
				lineNumber++
				out <- newCommand(lineNumber, line)
			}
		}
	}()
	return out
}

func isCommentOrEmptyLine(line *string) bool {
	return len(*line) == 0 || (*line)[0] == '/'
}
