package main

func solution(n int) int {
	result := 1
	k := 1

	for result <= n {
		k++
		result *= k
	}

	return k - 1
}