func solution(n int, numlist []int) []int {
    result := []int{}

    for i := 0; i < len(numlist); i++ {
        if numlist[i] % n == 0 {
            result = append(result, numlist[i])
        }
    }
    return result
}