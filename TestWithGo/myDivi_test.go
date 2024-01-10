package maint

import "testing"

func TestDivision(t *testing.T) {
	result := Division(6, 3)
	expected := 2

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
