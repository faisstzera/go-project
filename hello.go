package main

import "fmt"

const spanishLang = "Spanish"

const frenchLang = "French"

const englishHelloPrefix = "Hello, "

const spanishHelloPrefix = "Hola, "

const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanishLang:
		prefix = spanishHelloPrefix
	case frenchLang:
		prefix = frenchHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
