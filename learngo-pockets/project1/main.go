package main

import "fmt"

func main() {
	greeting := greet("en")
	fmt.Println(greeting)
}

type language string

var phrasebook = map[language]string{
	"en": "Hello, World",
	"el": "Χαίρετε Κόσμε",
	"fr": "Bonjour le monde",
	"ar": "اهلا بالعالم",
	"ur": "ہیلو دنیا",
	"vi": "Xin chào Thế Giới",
}

func greet(l language) string {
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}
	return greeting
}
