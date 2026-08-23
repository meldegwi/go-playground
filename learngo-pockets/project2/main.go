package main

import (
	"fmt"
	"os"
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

	for _, cbook := range cBooks {
		fmt.Println("-", cbook.Book.Title, "by", cbook.Book.Author)
		fmt.Printf("owned by: %s\n\n", cbook.Holders)
	}
}
