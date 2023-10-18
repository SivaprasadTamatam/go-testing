package benchmark

import "testing"

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
