func solution(balls int, share int) int {
	n := balls
	m := share

	if m > n-m {
		m = n - m
	}

	answer := 1

	for i := 0; i < m; i++ {
		answer *= (n - i)
		answer /= (i + 1)
	}

	return answer
}
