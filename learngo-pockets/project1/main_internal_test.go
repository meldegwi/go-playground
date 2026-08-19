package main

import "testing"

func Example_main() {
	main()
	// Output:
	// Hello, World!
}

func TestGreet_English(t *testing.T) {
	lang := language("en")
	want := "Hello, World!"

	got := greet(lang)

	if got != want {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}

func TestGreet_French(t *testing.T) {
	lang := language("fr")
	want := "Bonjour le monde!"

	got := greet(lang)

	if got != want {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}

func TestGreet_UnsupportedLanguage(t *testing.T) {
	lang := language("abc")
	want := "Unsupported language :( Please try again."

	got := greet(lang)

	if got != want {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}
