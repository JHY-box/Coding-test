func solution(n int, t int) int {
	answer := 0

	for i := 1; i <= t; i++ {
		n *= 2
		answer = n
	}

	return answer
}