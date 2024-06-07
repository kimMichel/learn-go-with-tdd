package structuremethodsinterfaces

import "testing"

func TestArea(t *testing.T) {
	testsArea := []struct {
		form Form
		expected float64
	} {
		{form: Rectangle{Width: 12,Height:  6}, expected: 72.0},
		{form: Circle{Radius: 10}, expected: 314.1592653589793},
		{form: Triangle{Width: 12,Height: 6}, expected: 36.0},
	}

	for _, tt := range testsArea {
		result := tt.form.Area()
		if result != tt.expected {
			t.Errorf("%#v expected: %.2f, result: %.2f",tt.form, tt.expected, result)
		}
	}
}