package maps

import (
	"testing"
)

func TestFind(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("a know word", func(t *testing.T) {
		result, _ := dictionary.Find("test")
		expect := "this is just a test"

		compareString(t, result, expect)
	})

	t.Run("a unknow word", func(t *testing.T) {
		_, result := dictionary.Find("unknow")

		compareError(t, result, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		dictionary.Add(word, definition)

		compareDefinition(t, dictionary, word, definition)
	})

	t.Run("existent word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		compareError(t, err, ErrExistentWord)
		compareDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existent word", func(t *testing.T) {
		word := "test"
		definition := "its just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		dictionary.Update(word, newDefinition)

		compareDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("not existent word", func(t *testing.T) {
		word := "test"
		definition := "its just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		compareError(t, err, ErrUnexistentWord)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "definition of test"}

	dictionary.Delete(word)

	_, err := dictionary.Find(word)
	if err != ErrNotFound {
		t.Errorf("expected %s was deleted", word)
	}
}

func compareError(t *testing.T, result, expect error) {
	t.Helper()

	if result != expect {
		t.Errorf("expected: %s, error result: %s", expect, result)
	}
}

func compareDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	result, err := dictionary.Find(word)
	if err != nil {
		t.Fatal("supposed to find the added word:", err)
	}
	
	if definition != result {
		t.Errorf("expected: %s, result: %s", definition, result)
	}
}