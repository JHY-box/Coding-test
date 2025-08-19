func solution(i int, j int, k int) int {
	count := 0
	for x := i; x <= j; x++ {
		if x == 0 {
			if k == 0 {
				count++
			}
			continue
		}
		n := x
		for n > 0 {
			if n%10 == k {
				count++
			}
			n /= 10
		}
	}
	return count
}