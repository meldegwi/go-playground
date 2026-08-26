package main

import (
	"bufio"
	"os"

	"gordle/gordle"
)

func main() {
	g := gordle.New(bufio.NewReader(os.Stdin))
	g.Play()
}
