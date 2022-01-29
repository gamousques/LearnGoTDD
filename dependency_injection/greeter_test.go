package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Gastón")

	got:= buffer.String()
	want := "Hello, Gastón"

	if got != want {
		t.Errorf("got: %q want: %q", got, want)
	}
}