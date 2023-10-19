package benchmark

import "testing"

func TestCalculateFibonacci(t *testing.T) {
	result := CalculateFibonacci(3)

	if result != 2 {
		t.Errorf("Expected 2, but received %d ", result)
	}
}

func TestSum(t *testing.T) {
	result := Sum(10, 20, 30)

	if result != 60 {
		t.Errorf("Expected  Sum(10,20,30) = 60, but got %d", result)
	}
}
func BenchmarkCalculateFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = CalculateFibonacci(10)
	}

}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := Sum(10, 20, 30)

		if result != 60 {
			b.Errorf("the sum of Sum(10,20,30) != 60, but we received : %d", result)
		}
	}
}
