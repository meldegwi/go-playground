package main

import (
	"encoding/json"
	"os"
)

// A Bookworm contains the list of books on a bookworm's shelf.
type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

// Book describes a book on a bookworm's shelf.
type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

// loadBookworms reads the file and returns the list of bookworms,
// and their 	beloved books, found therein
func loadBookworms(filePath string) ([]Bookworm, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var bookworms []Bookworm

	err = json.NewDecoder(f).
		Decode(&bookworms)
	if err != nil {
		return nil, err
	}

	return bookworms, nil
}

// findCommonBooks return struct contains books that are on more than
// one bookworm shelf along with its owners.
func findCommonBooks(bookworms []Bookworm) map[Book][]string {
	res := make(map[Book][]string)

	if len(bookworms) < 2 {
		return res
	}

	cbmap := make(map[Book][]string)
	for i := range bookworms {
		bw := &bookworms[i]
		for _, book := range bw.Books {
			cbmap[book] = append(cbmap[book], bw.Name)
		}
	}

	for book, elem := range cbmap {
		if len(elem) > 1 {
			res[book] = elem
		}
	}

	return res
}
