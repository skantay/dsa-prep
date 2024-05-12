package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return nil
    }
    dumb := head

    var last *ListNode

    var l int

    for dumb != nil {
        l++
        if dumb.Next == nil {
            last = dumb
        }
        dumb = dumb.Next
    }

    toRotate := k % l

    if toRotate == 0 {
        return head
    }

    last.Next = head
    
    dumb = head

    for i := 0; i < toRotate; {
        if dumb.Next == last {
            i++
            last = dumb

            continue
        }
        dumb = dumb.Next    
    }

    head = last.Next

    last.Next = nil

    return head
}
