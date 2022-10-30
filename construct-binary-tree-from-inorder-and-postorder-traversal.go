func buildTree(inorder []int, postorder []int) *TreeNode {
    if len(inorder) == 0 || len(postorder) == 0 {
        return nil
    }
    temp := make(map[int]int)
    for i, v := range inorder{
        temp[v] = i
    }
    var dfs func(int, int)*TreeNode
    dfs = func(left, right int)*TreeNode{
       if left > right{
           return nil
       }
        value := postorder[len(postorder)-1]
        postorder = postorder[:len(postorder)-1]
        pos := temp[value]
        root := &TreeNode{value, nil, nil}
        root.Right = dfs(pos+1, right)
        root.Left = dfs(left, pos-1)
        return root
    }
    return dfs(0, len(inorder)-1)
}