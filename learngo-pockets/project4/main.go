package main

import (
	"flag"
	"fmt"
	"os"

	"gordle/gordle"
)

func main() {
	var filePath string
	flag.StringVar(&filePath, "corpus-fp", "corpus/words",
		"Define corpus file pass as needed.")
	flag.Parse()

	corpus, err := gordle.ReadCorpus(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
	}

	g, err := gordle.New(os.Stdin, corpus, 5)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
	}

	g.Play()
}
