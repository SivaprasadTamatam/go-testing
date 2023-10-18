package benchmark

func CalculateFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return CalculateFibonacci(n-1) + CalculateFibonacci(n-2)
}

func Sum(a, b, c int) int {
	return a + b + c
}
