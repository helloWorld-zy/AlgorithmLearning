> Problem: [387. 字符串中的第一个唯一字符](https://leetcode.cn/problems/first-unique-character-in-a-string/description/)

[TOC]

# 思路
> 遍历字符串存哈希表

# 解题方法
> 遍历字符串存下标，若一个字符再次出现则将其下标直接设为数组长度（这个数随意，只要确保越界就行）；最终遍历哈希表，从没越界的下标里找最小

# 复杂度
- 时间复杂度: 
> $O(n)$

- 空间复杂度: 
> $O(n)$

# Code
```Go []

func firstUniqChar(s string) int {
    res := len(s)
    temp := map[byte]int{}
    for i, _ := range s{
        if _, ok := temp[s[i]];ok{
            temp[s[i]] = len(s)+1
        }else{
            temp[s[i]]=i
        }
    }
    if len(temp)==0{
        return -1
    }
    for _, v := range temp{
        if res > v{
            res = v
        }
    }
    if res ==len(s){
        return -1
    }
    return res
}
```