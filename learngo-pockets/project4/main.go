package main

import (
	"os"

	"gordle/gordle"
)

func main() {
	g := gordle.New(os.Stdin, "hello", 5)
	g.Play()
}
