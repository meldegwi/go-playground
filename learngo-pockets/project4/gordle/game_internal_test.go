package gordle

import (
	"errors"
	"slices"
	"strings"
	"testing"
)

func TestGameAsk(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []rune
	}{
		{
			name:  "5-chars-english",
			input: "Hello",
			want:  []rune("Hello"),
		},
		{
			name:  "5-chars-arabic",
			input: "مرحبا",
			want:  []rune("مرحبا"),
		},
		{
			name:  "5-chars-japanese",
			input: "こんにちは",
			want:  []rune("こんにちは"),
		},
		{
			name:  "3-chars-japanese",
			input: "こんに\nこんにちは",
			want:  []rune("こんにちは"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := New(strings.NewReader(tc.input))
			got := g.ask()
			if !slices.Equal(got, tc.want) {
				t.Errorf("got: %v, want: %v", string(got), string(tc.want))
			}
		})
	}
}

func TestValidateGuess(t *testing.T) {
	tests := []struct {
		name  string
		guess []rune
		want  error
	}{
		{
			name:  "valid guess length",
			guess: []rune("Hello"),
			want:  nil,
		},
		{
			name:  "too long",
			guess: []rune("Too Long"),
			want:  errInvalidWordLen,
		},
		{
			name:  "too short",
			guess: []rune("Hi"),
			want:  errInvalidWordLen,
		},
		{
			name:  "empty word",
			guess: []rune(""),
			want:  errInvalidWordLen,
		},
		{
			name:  "nil word",
			guess: nil,
			want:  errInvalidWordLen,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := New(nil)
			got := g.validateGuess(tc.guess)
			if !errors.Is(got, tc.want) {
				t.Errorf("word %c, got %q, want %q", tc.guess, got, tc.want)
			}
		})
	}
}
