func findNumberOfLIS(nums []int) int {
    if len(nums) <= 1 {
        return len(nums)
    }
    dp, count, max, ans := make([]int, len(nums)), make([]int, len(nums)), 1, 0
    dp[0], count[0] = 1, 1
    for i := 1; i < len(nums); i++ {
        dp[i], count[i]= 1, 1
        for j := i - 1; j >= 0; j-- {
            if nums[i] > nums[j] {
                if dp[j] + 1 == dp[i] {
                    count[i] += count[j]
                } else if dp[j] + 1 > dp[i] {
                    count[i] = count[j]
                    dp[i] = dp[j] + 1
                }
            }
            if dp[i] > max {
                max = dp[i]
            }
        }
    }
    for i := 0; i < len(dp); i++ {
        if dp[i] == max {
            ans += count[i]
        }
    }
    return ans
}
