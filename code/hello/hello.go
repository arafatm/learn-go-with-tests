package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bon jour, "

func Hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}

	if lang == "spanish" {
		return spanishHelloPrefix + name
	}

	if lang == "french" {
		return frenchHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("", ""))
}
