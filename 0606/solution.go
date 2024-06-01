package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func tree2str(root *TreeNode) string {
    if root == nil {
        return ""
    }

    if root.Left == nil && root.Right == nil{
        return fmt.Sprintf("%v", root.Val)
    } else if root.Right != nil && root.Left == nil {
        return fmt.Sprintf("%v()(%v)",root.Val, tree2str(root.Right))
    } else if root.Right == nil && root.Left != nil {
        return fmt.Sprintf("%v(%v)", root.Val, tree2str(root.Left))
    }

    return fmt.Sprintf("%v(%v)(%v)", root.Val, tree2str(root.Left), tree2str(root.Right))
}

