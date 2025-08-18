package main

import "math"

func solution(n int) int {
	answer := 0
	sqrtN := int(math.Sqrt(float64(n)))

	for i := 1; i <= sqrtN; i++ {
		if n%i == 0 {
			answer += 2
		}
	}

	if sqrtN*sqrtN == n {
		answer--
	}

	return answer
}