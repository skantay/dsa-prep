package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    slow := head
    fast := head

    isDomalak := true

    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next

        if slow == fast {
            isDomalak = false
            break
        }
    }

    if isDomalak {
        return nil
    }

    for head != slow {
        head = head.Next
        slow = slow.Next
    }

    return head
}
