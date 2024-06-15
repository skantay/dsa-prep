/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorderTraversal(root *TreeNode) []int {
    arr := make([]int, 0)
    rec(root, &arr)

    return arr
}

func rec(root *TreeNode, arr *[]int){
    if root == nil {
        return
    }

    rec(root.Left, arr)
    rec(root.Right,arr)
    *arr = append(*arr, root.Val)

    return 
}
