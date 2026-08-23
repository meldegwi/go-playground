package main

import (
	"bufio"
	"cmp"
	"encoding/json"
	"os"
	"slices"
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

// A CommonBook describes which bookworms have books in common.
type CommonBook struct {
	Book    Book
	Holders []string
}

// loadBookworms reads the file and returns the list of bookworms,
// and their 	beloved books, found therein
func loadBookworms(filePath string) ([]Bookworm, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buffedReader := bufio.NewReaderSize(f, 1024*1024)
	var bookworms []Bookworm
	err = json.NewDecoder(buffedReader).
		Decode(&bookworms)
	if err != nil {
		return nil, err
	}

	return bookworms, nil
}

// findCommonBooks return array of contains all common book(s),
// along with its holders sorted by book title.
func findCommonBooks(bookworms []Bookworm) []CommonBook {
	if len(bookworms) < 2 {
		return nil
	}

	cbmap := make(map[Book][]string)
	for i := range bookworms {
		bw := &bookworms[i]
		for _, book := range bw.Books {
			cbmap[book] = append(cbmap[book],
				bw.Name)
		}
	}

	var res []CommonBook
	for book, holders := range cbmap {
		if len(holders) > 1 {
			res = append(res,
				CommonBook{
					Book:    book,
					Holders: holders,
				})
		}
	}

	slices.SortFunc(res, func(a, b CommonBook) int {
		return cmp.Or(
			cmp.Compare(a.Book.Title, b.Book.Title),
		)
	})

	return res
}
