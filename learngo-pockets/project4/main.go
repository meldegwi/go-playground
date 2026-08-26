package main

import (
	"os"

	"gordle/gordle"
)

func main() {
	g := gordle.New(os.Stdin, "mdgwi", 5)
	g.Play()
}
