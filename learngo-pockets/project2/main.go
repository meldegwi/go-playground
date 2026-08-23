package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var filePath string
	flag.StringVar(&filePath,
		"filepath",
		"testdata/bookworms.json",
		"Specify the json file we want to parse")
	flag.Parse()

	bookworms, err := loadBookworms(filePath)
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
