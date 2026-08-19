package main

import "testing"

func Example_main() {
	main()
	// Output:
	// Hello, World!
}

func TestGreet(t *testing.T) {
	want := "Hello, World!"

	got := greet()

	if got != want {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}
