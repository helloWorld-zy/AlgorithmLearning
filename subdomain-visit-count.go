func subdomainVisits(cpdomains []string) []string {
    ans:=[]string{}
    storage := map[string]int{}
    firstSplit :=[]string{}
    for _, s := range cpdomains{
        firstSplit = strings.Split(s, " ")
        count,_ := strconv.Atoi(firstSplit[0])
        keywords:=strings.Split(firstSplit[1],".")
        fmt.Println(keywords)
        if len(keywords)==1{
            storage[keywords[0]]+=count
        }
        if len(keywords)==2{
            storage[keywords[1]]+=count
            storage[keywords[0]+"."+keywords[1]]+=count
        }
        if len(keywords)==3{
            storage[keywords[2]]+=count
            storage[keywords[1]+"."+keywords[2]]+=count
            storage[keywords[0]+"."+keywords[1]+"."+keywords[2]]+=count
        }
    }
    for k,v :=range storage{
        sentence:= strconv.Itoa(v)+" "+k
        ans = append(ans, sentence)
    }
    fmt.Println(storage)
    return ans
}