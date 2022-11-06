func shipWithinDays(weights []int, days int) int {
    left, right := 0, 0
    for _, weight := range weights{
        if weight > left{
            left = weight
        }
        right += weight
    }
    for left < right {
        countOfDays, capacity, mid := 1, 0, (left + right) / 2
        for _, weight := range weights{
            if capacity + weight > mid{
                countOfDays+=1
                capacity = 0
            }
            capacity += weight
        }
        if countOfDays <= days{
            right = mid
        }else{
            left = mid + 1
        }
    }
    return left
}