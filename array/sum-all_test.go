package array

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	result := SumAll([]int{1, 2, 3}, []int{1, 2, 3, 4, 5})
	expect := []int{6, 15}

	if !reflect.DeepEqual(result, expect) {
		t.Errorf("expected: '%v', result: '%v'", expect, result)
	}
}