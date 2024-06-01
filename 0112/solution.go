package leetcode

func hasPathSum(root *TreeNode, targetSum int) bool {
    _, ok := rec(root, targetSum, 0)

    return ok
}

func rec(root *TreeNode, targetSum, sum int) (int, bool) {
    if root == nil {
        return sum, false
    }

    sum += root.Val

    if sum == targetSum && root.Right == nil && root.Left == nil {
        return 0, true
    }

    _, ok := rec(root.Right, targetSum, sum)
    if ok {
       return 0, true 
    }
    _, ok2 := rec(root.Left, targetSum, sum)
    if ok2 {
       return 0, true 
    }

    return 0, false
}
