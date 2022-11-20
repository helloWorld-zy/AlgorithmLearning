func climbStairs(n int) int {
    if n<4{
        return n
    }
    minimum, maximum := 2,3
    for i:=4;i<=n;i++{
        maximum +=minimum
        minimum = maximum-minimum
    }
    return maximum
}