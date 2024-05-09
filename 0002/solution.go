package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    return addTwoNumbersHelper(l1, l2, 0)
}

func addTwoNumbersHelper(l1 *ListNode, l2 *ListNode, left int) *ListNode {
    if l1 == nil && l2 == nil && left == 0 {
        return nil
    }

    sum := left
    if l1 != nil {
        sum += l1.Val
        l1 = l1.Next
    }
    if l2 != nil {
        sum += l2.Val
        l2 = l2.Next
    }
    
    nextleft := sum / 10
    
    result := &ListNode{Val: sum % 10}
    
    result.Next = addTwoNumbersHelper(l1, l2, nextleft)
    
    return result
}

