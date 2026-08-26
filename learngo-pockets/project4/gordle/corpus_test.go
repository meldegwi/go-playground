package gordle

import (
	"errors"
	"testing"
)

func TestReadCorpus(t *testing.T) {
	tests := []struct {
		name   string
		file   string
		length int
		err    error
	}{
		{
			name:   "english corpus",
			file:   "../corpus/words",
			length: 100,
			err:    nil,
		},
		{
			name:   "empty corpus",
			file:   "../corpus/empty",
			length: 0,
			err:    ErrCorpusIsEmpty,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ReadCorpus(tc.file)

			if len(got) != tc.length {
				t.Errorf("file length unmatched. expected %d, got %d", tc.length, len(got))
			}

			if !errors.Is(err, tc.err) {
				t.Errorf("returned errors doesn't match. expected %v, got %v", tc.err, err)
			}
		})
	}
}
