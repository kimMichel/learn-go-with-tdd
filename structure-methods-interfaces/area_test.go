package structuremethodsinterfaces

import "testing"

func TestArea(t *testing.T) {

	checkArea := func (t *testing.T, form Form, expect float64)  {
		t.Helper()
		result := form.Area()

		if result != expect {
			t.Errorf("expected: %.2f, result: %.2f", expect, result)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		result := circle.Area()
		expect := 314.15926535897934

		if result != expect {
			t.Errorf("expected: %.2f result: %.2f", expect, result)
		}
	})
}