package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const solutionLength = 5

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
	guess := g.ask()

	fmt.Printf("Your guess is: %s\n", string(guess))
}

func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d-character guess:\n", solutionLength)

	for {
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Gordle failed to read your guess: %s\n",
				err.Error())
		}

		guess := []rune(string(playerInput))
		if len(guess) != solutionLength {
			fmt.Fprintf(os.Stderr,
				"Oh oh... Your guess is %d-character but it must be a %d-character word. Please try again.\n",
				len(guess), solutionLength)
		} else {
			return guess
		}
	}
}
