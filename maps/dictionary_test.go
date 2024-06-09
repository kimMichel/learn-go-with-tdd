package maps

import "testing"

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
	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")

	expect := "this is just a test"
	result, err := dictionary.Find("test")

	if err != nil {
		t.Fatal("not was possible to find the added word:", err)
	}

	if expect != result {
		t.Errorf("expected: %s, result: %s", expect, result)
	}
}

func compareError(t *testing.T, result, expect error) {
	t.Helper()

	if result != expect {
		t.Errorf("expected: %s, error result: %s", expect, result)
	}
}