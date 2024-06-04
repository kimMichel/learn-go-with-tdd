package main

import "testing"

func TestHello(t *testing.T) {
	result := Hello("Michel")
	expected := "Hello, Michel"

	if result != expected {
		t.Errorf("result: '%s', expected: '%s'", result, expected)
	}
}