package main

import (
	"bytes"
	"fmt"
)

func Greet(out *bytes.Buffer, name string) {
	fmt.Fprintf(out, "Hello, %s", name)
}