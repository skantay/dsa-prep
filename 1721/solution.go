package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapNodes(head *ListNode, k int) *ListNode {
    var l int

    dumb := head

    for ; dumb != nil; dumb = dumb.Next {
        l++
    }

    if l == 1 {
        return head
    }

    var swapF *ListNode

    var swapB *ListNode

    dumb = head

    for i := 0; i < l; i++ {
        if i == k-1 {
            swapF = dumb
        }
        if l - k == i {
            swapB = dumb
        }

        dumb = dumb.Next
    }
    
    swapF.Val, swapB.Val = swapB.Val, swapF.Val

    return head
}
