package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

var errInvalidWordLen = fmt.Errorf("invalid guess, word doesn't have the same number of characters as the solution")

// Game holds all the information we need to play a game of gordle.
type Game struct {
	reader      *bufio.Reader
	solution    []rune
	maxAttempts int
}

func New(pInput io.Reader, sol string, mxAtmpt int) *Game {
	g := &Game{
		reader:      bufio.NewReader(pInput),
		solution:    splitToUpperCaseCharacters(sol),
		maxAttempts: mxAtmpt,
	}

	return g
}

func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")

	for curAttempt := 1; curAttempt <= g.maxAttempts; curAttempt++ {
		guess := g.ask()

		if slices.Equal(guess, g.solution) {
			fmt.Printf("🎉 You won! You found it in %d guess(es)! The word was: %s.\n",
				curAttempt, string(g.solution))
			return
		} else {
			fmt.Printf("That was not it 😞.\n\n")
			fmt.Printf("Umm... I see that you're struggling... Here is a hint...\n\n")
			fmt.Printf("%s\n\n", giveHint(guess, g.solution))
		}
	}

	fmt.Printf("😞 You've lost! The solution was: %s. \n",
		string(g.solution))
}

func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d-character guess:\n", len(g.solution))

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
	if len(guess) != len(g.solution) {
		return fmt.Errorf(
			"expected %d, got %d %w",
			len(g.solution), len(guess), errInvalidWordLen)
	}
	return nil
}

func splitToUpperCaseCharacters(input string) []rune {
	return []rune(strings.ToUpper(input))
}

func giveHint(guess, solution []rune) feedback {
	if len(guess) != len(solution) {
		fmt.Fprintf(os.Stderr,
			"Internal error! Guess and solution have different lengths: %d vs %d",
			len(guess), len(solution))
		return nil
	}

	cm := map[rune]int{}
	for _, run := range solution {
		cm[run]++
	}

	fb := make(feedback, len(guess))
	for i, run := range guess {
		if run == solution[i] {
			fb[i] = correctPos
			cm[run]--
		} else {
			fb[i] = unknownChar
		}
	}

	for i, run := range guess {
		if fb[i] == unknownChar {
			if val, ok := cm[run]; ok && val > 0 {
				fb[i] = wrongPos
				cm[run]--
			} else {
				fb[i] = absentChar
			}
		}
	}

	return fb
}
