package main

import "fmt"

const messagePrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name =  "World"
	} 
	
	return messagePrefix + name
}

func main() {
	fmt.Println(Hello("Gastón"))
}
