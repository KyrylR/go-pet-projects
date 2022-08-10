package worker

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Result struct {
	Line    string
	LineNum int
	Path    string
}

type Results struct {
	Inner []Result
}

func NewResult(line, path string, lineNum int) Result {
	return Result{line, lineNum, path}
}

func FindInFile(path, find string) *Results {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error:", err)
		}
	}(file)

	results := Results{make([]Result, 0)}

	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), find) {
			r := NewResult(scanner.Text(), path, lineNum)
			results.Inner = append(results.Inner, r)
		}
		lineNum += 1
	}
	if len(results.Inner) == 0 {
		return nil
	} else {
		return &results
	}
}
