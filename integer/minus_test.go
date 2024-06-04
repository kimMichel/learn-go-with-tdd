package integer

import "testing"

func TestMinus(t *testing.T) {
	checkResult := func (t *testing.T, expected, minus int)  {
		t.Helper()
		if minus != expected {
			t.Errorf("expected: '%d', result: '%d'", expected, minus)
		}
	}
	
	t.Run("return subtraction of two integers", func(t *testing.T) {
		minus := Minus(3, 2)
		expected := 1

		checkResult(t, expected, minus)
	})
	
}