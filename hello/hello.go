package main

import (
	"fmt"
)

const spanish = "Spanish"
const french = "French"
const turkish = "Turkish"
const englishHelloPrefix = "Hello"
const spanishHelloPrefix = "Hola"
const frenchHelloPrefix = "Bonjour"
const turkishhHelloPrefix = "Merhaba"

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := greetingPrefix(language)

	return fmt.Sprintf("%s, %s!", prefix, name)
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case turkish:
		prefix = turkishhHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("User", "S"))
}
