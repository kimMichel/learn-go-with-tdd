package array

import "testing"

func TestSum(t *testing.T) {

	t.Run("collection of any size of numbers", func(t *testing.T) {
		numbers := []int{ 1, 2, 3, 4, 5}

		result := Sum(numbers)
		expect := 15

		if result != expect {
			t.Errorf("numbers: '%v', expected: '%v', result: '%v'", numbers, expect, result)
		}
	})

}