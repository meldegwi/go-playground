package gordle

import (
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
