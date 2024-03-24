package main

type node struct {
	val  int
	key  int
	next *node
}

type MyHashMap struct {
	size  int
	nodes []*node
}

func Constructor() MyHashMap {
	return MyHashMap{
		size:  100,
		nodes: make([]*node, 100),
	}
}

func (this *MyHashMap) Put(key int, value int) {
	var hash int

	if key != 0 {
		hash = this.size % key
	}

	node := &node{
		val: value,
		key: key,
	}

	curr := this.nodes[hash]

	if curr == nil {
		this.nodes[hash] = node

		return
	}

	for {
		if curr.key == key {
			curr.val = value

			return
		}

		if curr.next == nil {
			curr.next = node

			return
		}

		curr = curr.next
	}
}

func (this *MyHashMap) Get(key int) int {
	var hash int

	if key != 0 {
		hash = this.size % key
	}

	if this.nodes[hash] == nil {
		return -1
	}

	for curr := this.nodes[hash]; curr != nil; curr = curr.next {
		if curr.key == key {
			return curr.val
		}
	}

	return -1
}

func (this *MyHashMap) Remove(key int) {
	var hash int

	if key != 0 {
		hash = this.size % key
	}

	curr := this.nodes[hash]
	if curr != nil && curr.key == key {
		this.nodes[hash] = curr.next

		return
	}
	prev := curr
	curr = curr.next

	for {
		if curr == nil {
			return
		}

		if curr.key == key {
			prev.next = curr.next

			return
		}

		prev = curr
		curr = curr.next
	}
}
