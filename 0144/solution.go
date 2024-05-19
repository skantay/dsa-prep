package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func preorderTraversal(root *TreeNode) []int {
    res := make([]int, 0)
    return rec(root, res)
}

func rec(root *TreeNode, res []int) []int {
    if root == nil {
        return res
    }

    res = append(res, root.Val)

    res = rec(root.Left, res)
    res = rec(root.Right, res)

    return res
}
