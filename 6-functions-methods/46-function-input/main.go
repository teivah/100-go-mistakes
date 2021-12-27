package main

import (
	"bufio"
	"io"
	"os"
)

func countEmptyLinesInFile(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	// Handle file closure

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// ...
	}

	return 0, nil
}

func countEmptyLines(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		// ...
	}
	return 0, nil
}

func main() {
	file, err := os.Open("main.go")
	if err != nil {
		panic(err)
	}
	_, _ = countEmptyLines(file)
}
