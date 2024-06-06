package structuremethodsinterfaces

import "testing"

func TestArea(t *testing.T) {
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		result := rectangle.Area()
		expect := 72.0

		if result != expect {
			t.Errorf("expected: %.2f result: %.2f", expect, result)
		}
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