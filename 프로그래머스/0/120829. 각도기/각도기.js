function solution(angle) {
    
    if (90>angle) {
        return 1
    } else if (angle == 90) {
        return 2
    } else if (angle < 180) {
        return 3
    } else if (angle == 180) {
        return 4
    }
}