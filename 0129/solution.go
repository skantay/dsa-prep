package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    return rec(root, 0)
}

func rec(root *TreeNode, sum int) int {
    if root == nil {
        return 0
    }
    v := (sum*10)+root.Val
    if root.Left == nil && root.Right == nil {
        return v
    }
    sumLeft := rec(root.Left, v)
    sumRight := rec(root.Right, v)

    return sumLeft + sumRight
}
