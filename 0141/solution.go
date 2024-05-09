package leetcode

func hasCycle(head *ListNode) bool {
    m := make(map[*ListNode]struct{})

    for ; head != nil; head = head.Next {
        _, ok := m[head]
        if ok {
            return true
        }
        m[head] = struct{}{}
    } 

    return false
}
