import (
	"math"
)

func solution(n int) int {
	sqrt := math.Sqrt(float64(n))

	if sqrt == float64(int(sqrt)) {
		return 1
	} else {
		return 2
	}
}