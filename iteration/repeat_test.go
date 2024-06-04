package iteration

import "testing"

func TestRepeat(t *testing.T) {
	reps := Repeat("a")
	expect := "aaaaa"

	if reps != expect {
		t.Errorf("expected: '%s', result: '%s'", expect, reps)
	}
}