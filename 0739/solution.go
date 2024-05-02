package leetcode

func dailyTemperatures(temperatures []int) []int {
    stack := make([]int, 0, len(temperatures))

    result := make([]int, len(temperatures))
    
    for i := 0; i < len(temperatures); i++ {
        for len(stack) != 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
            result[stack[len(stack)-1]] = i - stack[len(stack)-1]

            stack = stack[:len(stack)-1]
        }

        stack = append(stack, i)
    }

    return result
}
