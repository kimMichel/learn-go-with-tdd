package di

import (
	"bytes"
	"testing"
)

func TestGreeting(t *testing.T) {
	buffer := bytes.Buffer{}
	Greeting(&buffer, "Chris")

	result := buffer.String()
	expect := "Hello Chris"

	if result != expect {
		t.Errorf("expected: %s, result: %s", expect, result)
	}
}