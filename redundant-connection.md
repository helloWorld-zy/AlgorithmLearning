> Problem: [684. 冗余连接](https://leetcode.cn/problems/redundant-connection/description/)

[TOC]

# 思路
> 并查集

# 解题方法
>  遍历 edges 中所有边，将边的两个节点加入到并查集进行合并; 若两个节点根节点不同，说明还未形成环，继续遍历; 若两个节点根节点相同，说明形成环，返回该边的两个节点

# 复杂度
- 时间复杂度: 
> $O(nlogn)$

- 空间复杂度: 
> $O(n)$

# Code
```Go []

func findRedundantConnection(edges [][]int) []int {
    us := make([]int, len(edges)+1)
    for i:=0;i<len(edges)+1;i++{
        us[i]=i
    }
    for _, edge := range edges{
        x, y := edge[0], edge[1]
        if Find(us, x) == Find(us, y){
            return []int{x, y}
        }else{
            UnionSet(us, x, y)
        }
    }
    return []int{}
}

func Find(us []int, x int) int{
    if us[x]==x{
        return x
    }
    return Find(us, us[x])
}

func UnionSet(us []int, x,y int){
    x, y = Find(us, x), Find(us, y)
    if x!=y{
        us[x]=y
    }
}
```