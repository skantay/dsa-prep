package leetcode 

type MinStack struct {
    topStack    int
    topMinVal   int
    stack       []int
    minVal      []int
}


func Constructor() MinStack {
    return MinStack{
        topStack: -1,
        topMinVal: -1,
        minVal: make([]int, 0),
        stack:  make([]int, 0),
    }
}


func (m *MinStack) Push(val int)  {
    m.stack = append(m.stack, val)
    m.topStack++
    
    if m.topMinVal == -1 {
        m.minVal = append(m.minVal, val)
        m.topMinVal++
    } else {
        if val <= m.minVal[m.topMinVal] {
            m.minVal = append(m.minVal, val)
            m.topMinVal++
        }
    }
}


func (m *MinStack) Pop()  {
    if m.topStack == -1 || m.topMinVal == -1  {
        return
    }
    if m.stack[m.topStack] == m.minVal[m.topMinVal] {
        m.minVal = m.minVal[:m.topMinVal]
        m.topMinVal--
    }

    m.stack = m.stack[:m.topStack]
    m.topStack--
}


func (m *MinStack) Top() int {
    if m.topStack == -1 {
        return 0
    }

    return m.stack[m.topStack]
}


func (m *MinStack) GetMin() int {
    if m.topMinVal == -1  {
        return 0
    }

    return m.minVal[m.topMinVal]
}
