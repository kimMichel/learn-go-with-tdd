package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}

	result := Search(dictionary, "test")
	expect := "this is just a test"

	compareString(t, result, expect)
}

func compareString(t *testing.T, result, expect string) {
	t.Helper()

	if result != expect {
		t.Errorf("given: %s, expected: %s, result: %s", "test", expect, result)
	}
}