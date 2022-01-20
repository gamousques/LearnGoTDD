package main

import "fmt"

const messagePrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		return "Hello, World"
	} else {
		return messagePrefix + name
	}
}

func main() {
	fmt.Println(Hello("Gastón"))
}
