func searchMatrix(matrix [][]int, target int) bool {
    res := false
    if target<matrix[0][0] || target> matrix[len(matrix)-1][len(matrix[0])-1]{
        return false
    }
    if target==matrix[0][0] || target == matrix[len(matrix)-1][len(matrix[0])-1]{
        return true
    }
    if target < matrix[len(matrix)/2][0]{
        res = searchMatrix(matrix [:len(matrix)/2], target)
    }
    if target > matrix[len(matrix)/2][len(matrix[0])-1]{
        res = searchMatrix(matrix [len(matrix)/2+1:], target)
    }
    if target > matrix[len(matrix)/2][0] && target < matrix[len(matrix)/2][len(matrix[0])-1]{
        for _, v :=range matrix[len(matrix)/2]{
            if v == target{
                return true
            }
        }
    }
    if target == matrix[len(matrix)/2][0] || target == matrix[len(matrix)/2][len(matrix[0])-1]{
        return true
    }
    return res
}