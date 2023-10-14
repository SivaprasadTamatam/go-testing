package mathutil

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Add(2, 3) = %d; want 5", result)
	}
}

func TestDivide(t *testing.T) {

	result, err := Divide(6, 0)

	if err == nil {
		t.Errorf("Expected an error bot got result: %d", result)
	}
}

var multiTests = []struct {
	a, b, expected int
}{
	{2, 3, 6},
	{5, 4, 20},
	{0, 7, 0},
}

func TestMultiply(t *testing.T) {
	for _, tt := range multiTests {
		result := Multiply(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Multiply(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
		}
	}

}
