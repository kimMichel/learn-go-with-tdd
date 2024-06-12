package sync

import "testing"

func TestCounter(t *testing.T) {
	t.Run("increment the counter 3 times and results the value of 3", func(t *testing.T) {
		counter := Counter{}
		counter.Increment()
		counter.Increment()
		counter.Increment()

		if counter.Value() != 3 {
			t.Errorf("result: %d, expected: %d", counter.Value(), 3)
		}
	})
}