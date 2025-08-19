func solution(numbers []int, direction string) []int {
    n := len(numbers)

    if direction == "right" {
        return append([]int{numbers[n-1]}, numbers[:n-1]...)
    } else {
        return append(numbers[1:], numbers[0])
    }
}