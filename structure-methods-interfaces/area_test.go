package structuremethodsinterfaces

import "testing"

func TestArea(t *testing.T) {
	testsArea := []struct {
		form Form
		expected float64
	} {
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range testsArea {
		result := tt.form.Area()
		if result != tt.expected {
			t.Errorf("expected: %.2f, result: %.2f", tt.expected, result)
		}
	}
}