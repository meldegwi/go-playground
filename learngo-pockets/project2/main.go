package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	bookworms, err := loadBookworms("testdata/bookworms.json")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to load bookworms: %s\n", err)
		os.Exit(1)
	}

	displayCommonBooks(bookworms)
}

func displayCommonBooks(bookworms []Bookworm) {
	cBooks := findCommonBooks(bookworms)
	fmt.Printf("Between all bookworms, those are the common books:\n\n")
	for _, cbook := range cBooks {
		fmt.Println("-", cbook.Book.Title, "by", cbook.Book.Author)
		fmt.Printf("Owned by: %s\n\n", strings.Join(cbook.Holders, ", "))
	}
}
