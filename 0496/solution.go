package leetcode

type maxStack struct {
	top   int
	array []int
}

func newMaxStack(cap int) *maxStack {
	return &maxStack{
		top: -1,
		array: make([]int, 0, cap),
	}
}

func (m *maxStack) topVal() int {
    if m.isEmpty() {
        panic("empty")
    }

	return m.array[m.top]
}

func (m *maxStack) push(n int) {
	m.top++
	m.array = append(m.array, n)
}

func (m *maxStack) pop() int {
        if m.isEmpty() {
		panic("empty")
	}

	val := m.array[m.top]
    m.array = m.array[:m.top]
	m.top--
	return val
}

func (m *maxStack) isEmpty() bool {
    return len(m.array) == 0
}


func nextGreaterElement(nums1 []int, nums2 []int) []int {
    stack := newMaxStack(len(nums1))

    hashMap := make(map[int]int, len(nums2))

    for i := len(nums2) - 1; i >= 0; i--{

        for !stack.isEmpty() && stack.topVal() < nums2[i] {
            stack.pop()
        }

        if stack.isEmpty() {
            hashMap[nums2[i]] = -1
        } else {
            hashMap[nums2[i]] = stack.topVal()
       }

        stack.push(nums2[i])
    }

    list := make([]int, 0, len(nums1))
    
    for i := range nums1 {
        list = append(list, hashMap[nums1[i]])
    }

    return list
}

