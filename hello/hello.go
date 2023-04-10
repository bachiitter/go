package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	if lang == "Spanish" {
		return spanishHelloPrefix + name
	}

	if lang == "French" {
		return frenchHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("", ""))
}
