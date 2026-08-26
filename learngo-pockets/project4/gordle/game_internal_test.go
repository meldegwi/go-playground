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
			want:  []rune("HELLO"),
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
			g := New(strings.NewReader(tc.input), string(tc.want), 0)
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
			g := New(nil, string("Hello"), 0)
			got := g.validateGuess(tc.guess)
			if !errors.Is(got, tc.want) {
				t.Errorf("word %c, got %q, want %q", tc.guess, got, tc.want)
			}
		})
	}
}

func TestGiveFeedback(t *testing.T) {
	tests := []struct {
		name           string
		guess          string
		solution       string
		wantedFeedback feedback
	}{
		{
			name:     "correct guess",
			guess:    "Hello",
			solution: "Hello",
			wantedFeedback: feedback{
				correctPos,
				correctPos,
				correctPos,
				correctPos,
				correctPos,
			},
		},
		{
			name:     "mixed feedback",
			guess:    "World",
			solution: "Hello",
			wantedFeedback: feedback{
				absentChar,
				wrongPos,
				absentChar,
				correctPos,
				absentChar,
			},
		},
		{
			name:     "correct positions and absent chars",
			guess:    "Hxllo",
			solution: "Hello",
			wantedFeedback: feedback{
				correctPos,
				absentChar,
				correctPos,
				correctPos,
				correctPos,
			},
		},
		{
			name:     "all chars correct but wrong positions",
			guess:    "olleH",
			solution: "Hello",
			wantedFeedback: feedback{
				wrongPos,
				wrongPos,
				correctPos,
				wrongPos,
				wrongPos,
			},
		},
		{
			name:     "mixed feedback",
			guess:    "Holol",
			solution: "Hello",
			wantedFeedback: feedback{
				correctPos,
				wrongPos,
				correctPos,
				absentChar,
				wrongPos,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotFeedback := giveHint([]rune(tc.guess), []rune(tc.solution))

			if !tc.wantedFeedback.Equal(gotFeedback) {
				t.Errorf("guess: %q, got the wrong feedback, wanted %v, got %v",
					tc.guess, tc.wantedFeedback, gotFeedback)
			}
		})
	}
}
