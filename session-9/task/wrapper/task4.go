package wrapper

import (
	"fmt"
	"os"
)

func OpenFile2(filename string) error {
	file, _ := os.Open(filename)
	// if err != nil {
	// 	return fmt.Errorf("failed to read file: %w", err)
	// }

	err1 := readFile(file)
	if err1 != nil {
		return fmt.Errorf("failed to read file: %w", err1)
	}
	return nil
}

func readFile(f *os.File) error { // Returns wrapped error is the reading the file fails.
	return fmt.Errorf("failed to open file: ")

}
