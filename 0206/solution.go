package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode

    if head == nil {
        return head
    }

    prev = nil

    for head.Next != nil {
        next := head.Next
        cur := head
        head.Next = prev
        prev = cur
        head = next
    }

    head.Next = prev

    return head
}
