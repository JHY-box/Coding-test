func yjh(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	if a < 0 {
		return -a
	}
	return a
}

func solution(numer1 int, denom1 int, numer2 int, denom2 int) []int {
	numer := numer1*denom2 + numer2*denom1
	denom := denom1 * denom2

	g := yjh(numer, denom)
	return []int{numer / g, denom / g}
}