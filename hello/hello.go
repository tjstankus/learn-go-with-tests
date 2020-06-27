package main

import "fmt"

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return languagePrefix(language) + name
}

func languagePrefix(language string) string {
	switch language {
	case "Spanish":
		return "Hola, "
	case "English":
		return "Hello, "
	default:
		return "Hello, "
	}
}

func main() {
	fmt.Println(Hello("world", "English"))
}
