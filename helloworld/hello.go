package main

import "fmt"

const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name =  "World"
	} 

	return GetLanguagePrefix(language ) + name
}

func GetLanguagePrefix(language string) string {
	switch language {
		case "Spanish":
			return spanishPrefix
		case "French":
			return frenchPrefix
	default:
		return englishPrefix
	}
}

func main() {
	fmt.Println(Hello("Gastón", ""))
}
