package gordle

import (
	"bufio"
	"fmt"
	"io"
)

// Game holds all the information we need to play a game of gordle.
type Game struct {
	reader *bufio.Reader
}

func New(pInput io.Reader) *Game {
	g := &Game{
		reader: bufio.NewReader(pInput),
	}

	return g
}

func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")

	fmt.Printf("Enter a guess:\n")
}
