package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string {"test": "this is a test"}

	got := Search(dictionary, "test")
	want := "this is a test"

	assertStrings(t, got, want)

}

func assertStrings(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("got: %s want:%s", got, want)
	}
}