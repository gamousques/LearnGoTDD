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

func TestAdd(t *testing.T) {
	t.Run("add a word+definition and get definition", func(t *testing.T) {

		dictionary := Dictionary{}
		word := "test"
		definition := "this is a test"
		err := dictionary.Add(word, definition)
	
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
		
	})

	t.Run("add existing word expect error", func(t *testing.T) {
	
		word := "test"
		definition := "this is a test"
		dictionary := Dictionary{word:  definition}
		err := dictionary.Add(word, definition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)

	})


}

func TestUpdate(t *testing.T) {
	t.Run("updata a definition and read updates definition", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "this is the new definition"
		err := dictionary.Update(word, newDefinition)
		
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("Update a new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)
		
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("given a word delete definition", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "this is a definition"}
		dictionary.Delete(word)

		_, err := dictionary.Search(word)

		if err != ErrWordNotFound {
			t.Errorf("Expected %q to be deleted", word)
		}



	})
}
func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got: %s want:%s", got, want)
	}
}

func assertDefinition(t testing.TB, d Dictionary, word, definition string) {
	t.Helper()
	
	got, err := d.Search(word)

	if err != nil {
		t.Fatal("should find added word", err)
	}

	if got != definition {
		t.Errorf("got: %q want: %q", got, definition)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}



