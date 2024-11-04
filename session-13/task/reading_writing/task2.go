package reading_writing

import (
	"bytes"
	"fmt"
	"os"
)

func countLines(fileName string) int {
	var count int
	file := readFile(fileName)
	lines := bytes.Split(file, []byte("\n"))

	// Omit the empty lines
	for _, line := range lines {
		if len(line) != 0 {
			count++
		}
	}

	return count
}
func readFile(fileName string) []byte {
	file, _ := os.ReadFile(fileName)
	return file
}
func RunTask2() {
	res := countLines("task/reading_writing/story.txt")
	fmt.Println("Total lines in file:", res)
}
