package tax

import "testing"

func Test(t *testing.T) {
	result := addTax(12, 3)
	expected := 26

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
