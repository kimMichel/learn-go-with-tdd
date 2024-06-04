package integer

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	checkResult := func (t *testing.T, expected, sum int)  {
		t.Helper()
		if sum != expected {
			t.Errorf("expected: '%d', result: '%d'", expected, sum)
		}
	}

	t.Run("return sum of the two integers", func(t *testing.T) {
		sum := Sum(2, 2)
		expected := 4

		checkResult(t, expected, sum)
	})

	t.Run("return sum of the two negative integers", func(t *testing.T) {
		sum := Sum(-2, -2)
		expected := -4

		checkResult(t, expected, sum)
	})
}

func ExampleSum() {
	sum := Sum(1, 5)
	fmt.Println(sum)
	// Output: 6
}