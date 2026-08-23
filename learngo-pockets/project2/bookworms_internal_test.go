package main

import (
	"reflect"
	"testing"
)

var (
	handmaidsTale    = Book{Author: "Margaret Atwood", Title: "The Handmaid's Tale"}
	oryxAndCrake     = Book{Author: "Margaret Atwood", Title: "Oryx and Crake"}
	theBellJar       = Book{Author: "Sylvia Plath", Title: "The Bell Jar"}
	mrsDalloway      = Book{Author: "Virginia Woolf", Title: "Mrs Dalloway"}
	nineteen84       = Book{Author: "George Orwell", Title: "1984"}
	janeEyre         = Book{Author: "Charlotte Brontë", Title: "Jane Eyre"}
	toTheLighthouse  = Book{Author: "Virginia Woolf", Title: "To the Lighthouse"}
	braveNewWorld    = Book{Author: "Aldous Huxley", Title: "Brave New World"}
	fahrenheit451    = Book{Author: "Ray Bradbury", Title: "Fahrenheit 451"}
	wutheringHeights = Book{Author: "Emily Brontë", Title: "Wuthering Heights"}
)

func TestLoadBookworms(t *testing.T) {
	tests := []struct {
		name          string
		bookwormsFile string
		want          []Bookworm
		wantErr       bool
	}{
		{
			name:          "file exists",
			bookwormsFile: "testdata/bookworms.json",
			want: []Bookworm{
				{
					Name:  "Fadi",
					Books: []Book{handmaidsTale, theBellJar, mrsDalloway, nineteen84},
				},
				{
					Name:  "Peggy",
					Books: []Book{oryxAndCrake, handmaidsTale, janeEyre, toTheLighthouse, braveNewWorld},
				},
				{
					Name:  "Samir",
					Books: []Book{nineteen84, braveNewWorld, oryxAndCrake, fahrenheit451},
				},
				{
					Name:  "Clara",
					Books: []Book{theBellJar, janeEyre, mrsDalloway, wutheringHeights},
				},
			},
			wantErr: false,
		},
		{
			name:          "file does not exist",
			bookwormsFile: "testdata/no_file_here.json",
			want:          nil,
			wantErr:       true,
		},
		{
			name:          "invalid json",
			bookwormsFile: "testdata/invalid.json",
			want:          nil,
			wantErr:       true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name,
			func(t *testing.T) {
				got, err := loadBookworms(tc.bookwormsFile)

				if err == nil && tc.wantErr {
					t.Fatalf("expected an error, got none.")
				}

				if err != nil && !tc.wantErr {
					t.Fatalf("expected no error, got one: %s", err.Error())
				}

				if !reflect.DeepEqual(got, tc.want) {
					t.Fatalf("different results. got: %v, want: %v", got, tc.want)
				}
			})
	}
}

func TestFindCommonBooks(t *testing.T) {
	tests := []struct {
		name      string
		bookworms []Bookworm
		want      []CommonBook
	}{
		{
			name: "there is common books",
			bookworms: []Bookworm{
				{
					Name:  "Fadi",
					Books: []Book{handmaidsTale, theBellJar, mrsDalloway, nineteen84},
				},
				{
					Name:  "Peggy",
					Books: []Book{oryxAndCrake, handmaidsTale, janeEyre, toTheLighthouse, braveNewWorld},
				},
				{
					Name:  "Samir",
					Books: []Book{nineteen84, braveNewWorld, oryxAndCrake, fahrenheit451},
				},
				{
					Name:  "Clara",
					Books: []Book{theBellJar, janeEyre, mrsDalloway, wutheringHeights},
				},
			},
			want: []CommonBook{
				{
					Book:    nineteen84,
					Holders: []string{"Fadi", "Samir"},
				},
				{
					Book:    braveNewWorld,
					Holders: []string{"Peggy", "Samir"},
				},
				{
					Book:    janeEyre,
					Holders: []string{"Peggy", "Clara"},
				},
				{
					Book:    mrsDalloway,
					Holders: []string{"Fadi", "Clara"},
				},
				{
					Book:    oryxAndCrake,
					Holders: []string{"Peggy", "Samir"},
				},
				{
					Book:    theBellJar,
					Holders: []string{"Fadi", "Clara"},
				},
				{
					Book:    handmaidsTale,
					Holders: []string{"Fadi", "Peggy"},
				},
			},
		},
		{
			name: "all bookworms have the same book",
			bookworms: []Bookworm{
				{
					Name:  "Fadi",
					Books: []Book{theBellJar},
				},
				{
					Name:  "Peggy",
					Books: []Book{theBellJar},
				},
				{
					Name:  "Samir",
					Books: []Book{theBellJar},
				},
				{
					Name:  "Clara",
					Books: []Book{theBellJar},
				},
			},
			want: []CommonBook{
				{
					Book:    theBellJar,
					Holders: []string{"Fadi", "Peggy", "Samir", "Clara"},
				},
			},
		},
		{
			name: "no common books",
			bookworms: []Bookworm{
				{
					Name:  "Samir",
					Books: []Book{nineteen84, braveNewWorld, oryxAndCrake, fahrenheit451},
				},
				{
					Name:  "Clara",
					Books: []Book{theBellJar, janeEyre, mrsDalloway, wutheringHeights},
				},
			},
			want: nil,
		},
		{
			name:      "no bookworms provided",
			bookworms: []Bookworm{},
			want:      nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name,
			func(t *testing.T) {
				got := findCommonBooks(tc.bookworms)
				if ok := reflect.DeepEqual(got, tc.want); !ok {
					t.Fatalf("expected: %#v, got: %#v", tc.want, got)
				}
			})
	}
}
