package main

import "fmt"

const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "
const french = "French"
const frenchHelloPrefix = "Bonjour, "
const englishHelloPrefix = "Hello, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greeting(language) + name
}

func greeting(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
