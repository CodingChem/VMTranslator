package main

import (
	"os"
	"path"
	"strings"

	"github.com/codingchem/VMTranslator/src/core"
)

func main() {
	args := os.Args
	input := args[1]
	if isDirectory(input) {
		processDirectory(input)
	} else {
		processFile(input)
	}
}

func processFile(inputFileName string) {
	baseFileName := strings.Split(path.Base(inputFileName), ".")[0]
	outputFileName := baseFileName + ".asm"

	cmds := core.ParseFile(baseFileName + ".vm")

	file := core.VMFile{
		Commands: cmds,
		Name:     baseFileName,
	}
	core.WriteCode([]core.VMFile{file}, outputFileName)
}

func processDirectory(inputDirectory string) {
	directoryName := path.Base(inputDirectory)
	outputFileName := directoryName + ".asm"
	files := GetFiles(inputDirectory)
	vmFiles := []core.VMFile{}

	for _, file := range files {
		if strings.HasSuffix(file, ".vm") {
			baseFileName := strings.Split(file, ".")[0]
			cmds := core.ParseFile(inputDirectory + "/" + file)
			file := core.VMFile{
				Commands: cmds,
				Name:     baseFileName,
			}
			vmFiles = append(vmFiles, file)
		}
	}
	core.WriteCode(vmFiles, outputFileName)
}

func isDirectory(input string) bool {
	fileInfo, err := os.Stat(input)
	if err != nil {
		panic(err)
	}
	return fileInfo.IsDir()
}

func GetFiles(directory string) []string {
	allFiles, err := os.ReadDir(directory)
	if err != nil {
		panic(err)
	}
	files := []string{}
	for _, file := range allFiles {
		if strings.HasSuffix(file.Name(), ".vm") {
			files = append(files, file.Name())
		}
	}

	return files
}
