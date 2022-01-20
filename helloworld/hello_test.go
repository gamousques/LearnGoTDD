package main

import "testing"

func TestHello_returnOK(t *testing.T) {
	want := "Hello, Gastón"
	got := Hello("Gastón")

	if got != want {
		t.Errorf("got: '%s' want: '%s'", got, want)
	}
}
