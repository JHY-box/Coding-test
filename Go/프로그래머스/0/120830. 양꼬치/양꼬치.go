func solution(n int, k int) int {
	answer := 0
	answer = (n * 12000) + ((k - (n / 10)) * 2000)
	return answer
}