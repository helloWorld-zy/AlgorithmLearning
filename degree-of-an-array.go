func findShortestSubArray(nums []int) int {
    if len(nums)==1{
        return 1
    }
    // 变量按顺序分别是度， 每个数字的计数， 该数字的索引， 计数与度相同的数字， 该数字的子数组的长度
    degree, countOfEachNum, index, maxNums,length := 0, map[int]int{}, map[int][]int{}, []int{}, 0
    for i:=0;i<len(nums);i++{
        countOfEachNum[nums[i]]+=1
        index[nums[i]]=append(index[nums[i]],i)
        if degree<countOfEachNum[nums[i]]{
            degree = countOfEachNum[nums[i]]
        }
    }
    // 遍历计数数组，获取与度相等的数字
    for k,v := range countOfEachNum{
        if v==degree{
            maxNums = append(maxNums, k)
        }
    }
    // 如果与度相等的只有一个数字，则直接返回该数字所组成的连续子数组
    if len(maxNums)==1{
        return index[maxNums[0]][len(index[maxNums[0]])-1]-index[maxNums[0]][0]+1
    }
    // 遍历一遍index数组，获取最小长度
    for _, x := range maxNums{
        if length > (index[x][len(index[x])-1]-index[x][0]+1) || length==0{      
            length = index[x][len(index[x])-1]-index[x][0]+1
        }
    }
    return length
}