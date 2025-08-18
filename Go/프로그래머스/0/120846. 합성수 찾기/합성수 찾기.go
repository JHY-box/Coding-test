func solution(n int) int {
	count := 0

	for i := 1; i <= n; i++ {
		divisors := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				divisors++
			}
		}
		if divisors >= 3 {
			count++
		}
	}
	return count
}