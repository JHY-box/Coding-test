package main

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func solution(n int) int {
	lcm := n * 6 / gcd(n, 6)
	return lcm / 6
}