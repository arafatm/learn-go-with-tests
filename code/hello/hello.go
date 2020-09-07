package learngowithtests

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bon jour, "

// Hello World
func Hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}

	prefix := englishHelloPrefix

	switch lang {
	case "french":
		prefix = frenchHelloPrefix
	case "spanish":
		prefix = spanishHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("", ""))
}
