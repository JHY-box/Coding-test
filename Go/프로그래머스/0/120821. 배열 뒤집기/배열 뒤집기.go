func solution(num_list []int) []int {
    result := []int{}

    for i := len(num_list) - 1; i >= 0; i-- {
        result = append(result, num_list[i])
    }

    return result
}