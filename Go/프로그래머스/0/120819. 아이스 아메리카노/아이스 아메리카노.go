func solution(money int) []int {
    price := 5500
    cups := money / price
    change := money % price
    return []int{cups, change}
}