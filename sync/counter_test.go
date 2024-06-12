package sync

import "testing"

func TestCounter(t *testing.T) {
	t.Run("increment the counter 3 times and results the value of 3", func(t *testing.T) {
		counter := Counter{}
		counter.Increment()
		counter.Increment()
		counter.Increment()

		checkCounter(t, counter, 3)
	})
}

func checkCounter(t *testing.T, result Counter, expect int) {
	t.Helper()
	if result.Value() != expect {
		t.Errorf("result: %d, expected: %d", result.Value(), expect)
	}
}