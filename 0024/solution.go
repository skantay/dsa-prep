package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
   d, _ := rec(head, 0)

   return d
}

func rec(node *ListNode, cnt int) (*ListNode, int) {
    if node == nil {
        return nil, cnt
    }
    
    cnt++
    
    head, i := rec(node.Next, cnt)

    if i % 2 != 0 && head != nil {
        node.Next = head.Next
        head.Next = node

        return head, cnt - 1
    }

    node.Next = head

    return node, cnt - 1
}
