package lists

import "fmt"

type LinkedList struct {
	head *node
}

type node struct {
	val  any
	next *node
}

func New(val any) *LinkedList {
	return &LinkedList{
		head: &node{
			val: val,
		},
	}
}

func (l *LinkedList) Push(val any) {
	current := l.head

	for current.next != nil {
		current = current.next
	}

	current.next = &node{
		val: val,
	}
}

func (l *LinkedList) PushAt(val any, pos int) {
	if pos == 0 {
		node := &node{
			val:  val,
			next: l.head,
		}

		l.head = node

		return
	}

	node := &node{
		val: val,
	}

	current := l.head

	prev := current

	for i := 0; current != nil; i++ {
		if i == pos {
			prev.next = node
			node.next = current

			return
		}

		prev = current
		current = current.next
	}
}

// IsExists check if linked list have val or not
func (l *LinkedList) IsExists(val any) bool {
	current := l.head

	for current != nil {

		if val == current.val {
			return true
		}

		current = current.next
	}

	return false
}

// RemoveLast element from linked list (tail)
func (l *LinkedList) RemoveLast() {
	if l.head == nil {
		return
	}

	current := l.head

	prev := current

	var i int

	for ; current.next != nil; i++ {
		prev = current
		current = current.next
	}

	if i == 0 {
		l.head = nil
	} else {
		prev.next = nil
	}
}

// RemoveFirst element from linked list (head)
func (l *LinkedList) RemoveFirst() {
	if l.head != nil {
		l.head = l.head.next
	}
}

func (l *LinkedList) Print() {
	current := l.head

	for current != nil {

		if current.next == nil {
			fmt.Printf("%v", current.val)

			break
		}

		fmt.Printf("%v -> ", current.val)

		current = current.next
	}

	fmt.Println()
}
