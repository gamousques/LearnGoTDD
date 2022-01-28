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
	t.Helper()
	if got != want {
		t.Errorf("got: %s want:%s", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func TestAdd(t *testing.T) {
	t.Run("add a word+definition and get definition", func(t *testing.T) {

		dictionary := Dictionary{}
		word := "test"
		definition := "this is a test"
		err := dictionary.Add(word, definition)
	
		assertError(t, err, nil)
		assertDefition(t, dictionary, word, definition)
		
	})

	t.Run("add existing word expect error", func(t *testing.T) {
	
		word := "test"
		definition := "this is a test"
		dictionary := Dictionary{word:  definition}
		err := dictionary.Add(word, definition)
		assertError(t, err, ErrWordExists)
		assertDefition(t, dictionary, word, definition)

	})


}

func assertDefition(t testing.TB, d Dictionary, word, definition string) {
	t.Helper()
	
	got, err := d.Search(word)

	if err != nil {
		t.Fatal("should find added word", err)
	}

	if got != definition {
		t.Errorf("got: %q want: %q", got, definition)
	}
}

