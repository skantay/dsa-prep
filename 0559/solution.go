package leetcode

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
    if root == nil {
        return 0
    }
    
    return rec(root, 1)
}

func rec(root *Node, depth int) int {
    if root == nil {
        return depth
    }

    var m int = depth

    for _, next := range root.Children {
        m = max(m, rec(next, depth+1))
    }

    return m
}
