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

	fmt.Printf("%s\n", bookworms)
}
