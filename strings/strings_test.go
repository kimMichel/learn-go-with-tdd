package strings

import (
	"reflect"
	"testing"
)

func TestSplitText(t *testing.T) {
	parts := SplitText("Learn Go with Tests.")
	expect := []string{"Learn", "Go", "with", "Tests."}

	if !reflect.DeepEqual(parts, expect) {
		t.Errorf("expected: '%v', result: '%v'", expect, parts)
	}
}