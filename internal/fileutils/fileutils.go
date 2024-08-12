// Package fileutils provides utility functions for file operations.
package fileutils

import (
	"bufio"
	"os"
)

// ReadFileLines reads a file line by line and returns a channel from which lines can be read.
// It opens the file specified by the path and starts a goroutine to read the file.
// If there is an error opening the file, it returns the error.
func ReadFileLines(path string) (chan string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	channel := make(chan string)
	go func() {
		defer file.Close()
		defer close(channel)
		for scanner.Scan() {
			channel <- scanner.Text()
		}
	}()
	return channel, nil
}

// WriteFileLines writes lines to a file from a channel.
// If there is an error creating the file, it returns the error.
func WriteFileLines(path string, ch chan string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(file)
	defer func() {
		writer.Flush()
		file.Close()
	}()
	for line := range ch {
		writer.WriteString(line)
		writer.WriteRune('\n')
	}
	return nil
}
