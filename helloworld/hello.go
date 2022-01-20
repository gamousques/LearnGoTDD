package main

import "fmt"

const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name =  "World"
	} 

	var languagePrefix string

	switch language {
		case "Spanish":
			languagePrefix = spanishPrefix
		case "French":
			languagePrefix = frenchPrefix
	default:
		languagePrefix = englishPrefix
	}

	return languagePrefix + name
}

func main() {
	fmt.Println(Hello("Gastón", ""))
}
