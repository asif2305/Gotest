// example_test.go

package maint

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	expected := 2

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(5, 3)
	expected := 15

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
