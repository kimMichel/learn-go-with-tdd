package structuremethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	result := Perimeter(rectangle)
	expect := 40.0

	if result != expect {
		t.Errorf("expected: %.2f, result: %.2f", expect, result)
	}
}