package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	var i int

	h := head

	for head != nil {
		i++
		head = head.Next
	}

	head = h

	for j := 0; j <= i; j++ {

		if j == i/2 {
			return head
		}

		head = head.Next
	}

	return h
}
