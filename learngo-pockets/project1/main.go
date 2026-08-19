package main

import "fmt"

func main() {
	greeting := greet("en")
	fmt.Println(greeting)
}

type language string

func greet(l language) string {
	switch l {
	case "en":
		return "Hello, World!"
	case "fr":
		return "Bonjour le monde!"
	default:
		return "Unsupported language :( Please try again."
	}
}
