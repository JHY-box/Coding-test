func solution(num1 int, num2 int) int {
	answer := 0

	if num1 > 0 && num1 <= 100 && num2 > 0 && num2 <= 100 {
		num3 := (float64(num1) / float64(num2)) * 1000
		answer = int(num3)
	}

	return answer
}