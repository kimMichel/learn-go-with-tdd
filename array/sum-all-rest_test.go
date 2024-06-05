package array

import (
	"reflect"
	"testing"
)

func TestSumAllRest(t *testing.T) {
	
	checkSums := func (t *testing.T, expect, result []int)  {
		t.Helper()
		if !reflect.DeepEqual(result, expect) {
			t.Errorf("expect: '%v', result: '%v'", expect, result)
		}
	}

	t.Run("sum everything from index 1 to the right with 2 numbers on slice", func(t *testing.T) {
		result := SumAllRest([]int{1, 2}, []int{0, 9})
		expect := []int{2, 9}

		checkSums(t, expect, result)
	})

	t.Run("sum empty slices", func(t *testing.T) {
		result := SumAllRest([]int{}, []int{3, 4, 5})
		expect := []int{0, 9}

		checkSums(t, expect, result)
	})

}