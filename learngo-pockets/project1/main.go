package main

import (
	"flag"
	"fmt"
)

func main() {
	var lang string
	flag.StringVar(&lang, "lang",
		"en",
		"Use the flag to specify the language e.g. en, fr, ar etc.")
	flag.Parse()

	greeting := greet(language(lang))
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
