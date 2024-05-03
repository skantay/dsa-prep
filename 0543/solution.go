package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    res := 0
    delai(root, &res)
    return res
}


func delai(root *TreeNode, ans *int) int {
    if root == nil {
        return 0
    }

    l := delai(root.Left, ans)
    r := delai(root.Right, ans)

    *ans = max(*ans, l + r)

    return 1 + max(l, r)
}
