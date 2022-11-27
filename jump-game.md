> Problem: [55. 跳跃游戏](https://leetcode.cn/problems/jump-game/description/)

[TOC]

# 思路
> 用动态规划思路解

# 解题方法
> nums的len为1则必为true，直接返回；dp[i]代表下标i处最大跳跃值，该值等于dp[i-1]跳到dp[i]后剩余步数及nums[i]之间更大的拿一个，所以状态转移方程是dp[i] = max(dp[i-1], nums[i]);

# 复杂度
- 时间复杂度: 
> $O(n)$

- 空间复杂度: 
> $O(1)$

# Code
```Go []

func canJump(nums []int) bool {
    if len(nums)==1{
        return true
    }
    max := func(num1, num2 int)int{
        if num1>=num2{
            return num1
        }
        return num2
    }
    maxStep := nums[0]
    for i:=1;i<len(nums);i++{
        if maxStep == 0{
            return false
        }
        maxStep = max(maxStep-1, nums[i])
    }
    return true
}
```