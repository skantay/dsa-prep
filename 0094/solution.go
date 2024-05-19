package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    arr := make([]int, 0)
    rec(root, &arr)

    return arr
}


func rec(root *TreeNode, arr *[]int){
    if root == nil {
        return
    }

    rec(root.Left, arr)
    *arr = append(*arr, root.Val)
    rec(root.Right,arr)

    return 
}
