package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// Simple
func isPalindrome1(head *ListNode) bool {
	q := make([]int, 0)

	for ; head != nil; head = head.Next {
		q = append(q, head.Val)
	}

	var l, r int = 0, len(q) - 1

	for l < len(q) && r >= 0 {
		if q[l] != q[r] {
			return false
		}

		l++
		r--
	}

	return true
}

// Recursion
func isPalindrome2(head *ListNode) bool {
	_, ok := recursion(head, head)

	return ok
}

func recursion(head, front *ListNode) (*ListNode, bool) {
	if head == nil {
		return front, true
	}

	front, ok := recursion(head.Next, front)
	if !ok {
		return nil, false
	}

	if front == nil || head.Val != front.Val {
		return nil, false
	}

	return front.Next, true
}
