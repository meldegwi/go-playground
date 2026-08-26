package main

import (
	"os"

	"gordle/gordle"
)

func main() {
	g := gordle.New(os.Stdin, "mrDigo", 5)
	g.Play()
}
