package fileutils

import (
	"bufio"
	"os"
	"testing"
)

func TestReadFileLines(t *testing.T) {
	// Create a temporary file
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	// Write some lines to the file
	text := []byte("line1\nline2\nline3\n")
	if _, err := tmpfile.Write(text); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Use your function to read the lines
	lines, err := ReadFileLines(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}

	// Check that the lines are correct
	expectedLines := []string{"line1", "line2", "line3"}
	for i, line := range expectedLines {
		if actualLine := <-lines; actualLine != line {
			t.Errorf("Line %d: expected %q, got %q", i, line, actualLine)
		}
	}

	// Check that there are no more lines
	if _, more := <-lines; more {
		t.Error("Expected no more lines, but there were more")
	}
}

func TestWriteFileLines(t *testing.T) {
	// Create a temporary file
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	// Create a channel and write some lines to it
	lines := make(chan string)
	go func() {
		lines <- "line1"
		lines <- "line2"
		lines <- "line3"
		close(lines)
	}()

	// Use your function to write the lines to the file
	if err := WriteFileLines(tmpfile.Name(), lines); err != nil {
		t.Fatal(err)
	}

	// Read the file back and check the lines
	file, err := os.Open(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	expectedLines := []string{"line1", "line2", "line3"}
	for i, line := range expectedLines {
		if !scanner.Scan() {
			t.Fatal("Expected more lines")
		}
		if scanner.Text() != line {
			t.Errorf("Line %d: expected %q, got %q", i, line, scanner.Text())
		}
	}

	// Check that there are no more lines
	if scanner.Scan() {
		t.Error("Expected no more lines, but there were more")
	}
}
