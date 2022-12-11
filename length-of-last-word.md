> Problem: [58. 最后一个单词的长度](https://leetcode.cn/problems/length-of-last-word/description/)

[TOC]

# 思路
> 分割字符串，获得最后一个len不为0的字符串长度

# 复杂度
- 时间复杂度: 
> $O(n)$

- 空间复杂度: 
> $O(n)$

# Code
```Go []

func lengthOfLastWord(s string) int {
    var length int
    tempList := strings.Split(s," ")
    for i:=len(tempList)-1;i>=0;i--{
        if len(tempList[i])!=0{
            length = len(tempList[i])
            break
        }
    }
    return length
}
```