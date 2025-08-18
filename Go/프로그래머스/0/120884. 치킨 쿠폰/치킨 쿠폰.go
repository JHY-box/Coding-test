func solution(chicken int) int {
	answer := 0

	for chicken >= 10 {
		service := chicken / 10
		reminder := chicken % 10

		answer += service
		chicken = service + reminder
	}

	return answer
}