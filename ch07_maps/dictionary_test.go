package ch07_maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	word := "test"
	want := "this is just a test"
	assertDefinition(t, dictionary, word, want)
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := make(Dictionary, 1)
		word := "test"
		definition := "this is just a test"
		err := dictionary.Add(word, definition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("Existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, definition)
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"
		err := dictionary.Update(word, newDefinition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("New word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	dictionary := Dictionary{word: definition}
	assertDefinition(t, dictionary, word, definition)
	dictionary.Delete(word)
	_, err := dictionary.Search(word)
	assertError(t, err, ErrNotFound)
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, want string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("expected to get an error")
	}
	if got != want {
		t.Logf("got error %q want %q", got, want)
	}
}
