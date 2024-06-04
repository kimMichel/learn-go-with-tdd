package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	checkResults := func (t *testing.T, expect, reps string)  {
		t.Helper()
		if reps != expect {
			t.Errorf("expected: '%s', result: '%s'", expect, reps)
		}
	}
	
	t.Run("return the character 5 times", func(t *testing.T) {
		reps := Repeat("a", 5)
		expect := "aaaaa"

		checkResults(t, expect, reps)
	})
	
	t.Run("return the character according to the number of times passed in the parameter", func(t *testing.T) {
		reps := Repeat("a", 7)
		expect := "aaaaaaa"

		checkResults(t, expect, reps)
	})
}

func ExampleRepeat() {
	reps := Repeat("a", 5)
	fmt.Println(reps)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}