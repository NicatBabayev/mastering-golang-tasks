package structs

import (
	"fmt"
)

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b Book) String() string {
	return fmt.Sprintf(" Title: %s\n Author: %s\n Pages: %d\n", b.Title, b.Author, b.Pages)
}
