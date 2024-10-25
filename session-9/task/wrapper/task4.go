package wrapper

import (
	"fmt"
	"os"
)

func OpenFile2(filename string) error {
	_, err := os.Open(filename)
	if err != nil {
		err = fmt.Errorf("failed to read file: %w", err)
		return readFile(err)
	}
	return nil
}

func readFile(openErr error) error {

	return fmt.Errorf("failed to open file: %w", openErr)
}
