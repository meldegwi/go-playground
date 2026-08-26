package gordle

import (
	"fmt"
	"os"
	"strings"
)

const ErrCorpusIsEmpty = corpusError("corpus is empty")

func ReadCorpus(filePath string) ([]string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to open %q for reading: %w", filePath, err)
	}

	if len(data) == 0 {
		return nil, ErrCorpusIsEmpty
	}

	words := strings.Fields(string(data))

	return words, nil
}
