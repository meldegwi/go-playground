package main

import "testing"

func Example_main() {
	main()
	// Output:
	// Hello, World
}

func TestGreet(t *testing.T) {
	type testCases struct {
		lang language
		want string
	}

	tests := map[string]testCases{
		"English": {
			"en",
			"Hello, World",
		},
		"French": {
			"fr",
			"Bonjour le monde",
		},
		"Arabic": {
			"ar",
			"اهلا بالعالم",
		},
		"Unsupported Language": {
			"abc",
			`unsupported language: "abc"`,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := greet(tc.lang)

			if got != tc.want {
				t.Errorf("expected: %q, got: %q", tc.want, got)
			}
		})
	}
}
