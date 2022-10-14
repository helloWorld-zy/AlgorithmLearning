func plusOne(digits []int) []int {
    zeroCount:=0
    if len(digits)==1&&digits[0]!=9{
        digits[0]+=1
        return digits
    }else{
        for i:=len(digits)-1;i>=0;i--{
            if i==len(digits)-1&&digits[i]!=9{
                digits[i]+=1
                break
            }
            if i!=len(digits)-1&&digits[i]!=9{
                digits[i]++
                for j:=i+1;j<len(digits);j++{
                    digits[j]=0
                }
                break
            }
            digits[i]=0
            zeroCount++
        }
        if zeroCount==len(digits){
            digits = append([]int{1},digits...)
        }
    }
    return digits
}