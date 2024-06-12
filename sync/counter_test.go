package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("increment the counter 3 times and results the value of 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Increment()
		counter.Increment()
		counter.Increment()

		checkCounter(t, counter, 3)
	})

	t.Run("run on concurrency safety", func(t *testing.T) {
		expectedCounter := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(expectedCounter)

		for i := 0; i < expectedCounter; i++ {
			go func(w *sync.WaitGroup) {
				counter.Increment()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		checkCounter(t, counter, expectedCounter)
	})
}

func NewCounter() *Counter {
	return &Counter{}
}

func checkCounter(t *testing.T, result *Counter, expect int) {
	t.Helper()
	if result.Value() != expect {
		t.Errorf("result: %d, expected: %d", result.Value(), expect)
	}
}