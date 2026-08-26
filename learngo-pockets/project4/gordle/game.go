package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const solutionLength = 5

var errInvalidWordLen = fmt.Errorf("invalid guess, word doesn't have the same number of characters as the solution")

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

		guess := splitToUpperCaseCharacters(string(playerInput))
		err = g.validateGuess(guess)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Your attempt is invalid with Gordle solution: %s.\n", err.Error())
		} else {
			return guess
		}
	}
}

func (g *Game) validateGuess(guess []rune) error {
	if len(guess) != solutionLength {
		return fmt.Errorf(
			"expected %d, got %d %w",
			solutionLength, len(guess), errInvalidWordLen)
	}
	return nil
}

func splitToUpperCaseCharacters(input string) []rune {
	return []rune(strings.ToUpper(input))
}
