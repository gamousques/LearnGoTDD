package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	t.Run("given a word to search returns the correct value", func(t *testing.T) {
		

		got, _ := dictionary.Search( "test")
		want := "this is a test"
	
		assertStrings(t, got, want)
	})

	t.Run("given a non existent word returns an error", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrWordNotFound
	
		assertError(t, err, want)
	})

}

func assertStrings(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("got: %s want:%s", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}