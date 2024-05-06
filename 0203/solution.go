package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements1(head *ListNode, val int) *ListNode {
	var h, prev *ListNode = nil, head

	var foundHead bool

	for head != nil {
		if head.Val == val {
			prev.Next = head.Next
			head = head.Next

			continue
		} else if !foundHead {
			h = head
			foundHead = true
		}

		prev = head
		head = head.Next
	}

	return h
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements2(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	head.Next = removeElements2(head.Next, val)

	if head.Val == val {
		return head.Next
	}

	return head
}
