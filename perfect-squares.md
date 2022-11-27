> Problem: [279. 完全平方数](https://leetcode.cn/problems/perfect-squares/description/)

[TOC]

# 思路
> 根据作业要求选用背包问题解决

# 解题方法
> 和为n的背包所需最少完全平方数为dp[i]；递推公式为dp[i] = min(dp[i],dp[i-num]+1);

# 复杂度
- 时间复杂度: 
> $O(nlogn)$

- 空间复杂度: 
> $O(n)$

# Code
```Go []

func numSquares(n int) int {
    dp := make([]int, n+1)
    dp[0],dp[1]=0,1
    for i:=2;i<n+1;i++{
        dp[i]=n
    }
    isSqrt := func(num int)bool{
        numString :=strconv.Itoa(num)
        numFloat,_ := strconv.ParseFloat(numString,64)
        squareNum := math.Sqrt(numFloat)
        return squareNum-math.Trunc(squareNum)==0
    }
    min := func(num1, num2 int)int{
        if num1<=num2{
            return num1
        }
        return num2
    }
    for num := 0;num<n+1;num++{
        if isSqrt(num){
            for i:=num;i<n+1;i++{
                dp[i] = min(dp[i],dp[i-num]+1)
            }
        }
    }
    return dp[n]
}
```