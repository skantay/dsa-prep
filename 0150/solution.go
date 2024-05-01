package leetcode

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	var top int = -1

	for _, char := range tokens {
		switch char {
		case "+":
			num := stack[top]
			stack = stack[:top]
			top--
			stack[top] = stack[top] + num
		case "-":
			num := stack[top]
			stack = stack[:top]
			top--
			stack[top] = stack[top] - num
		case "*":
			num := stack[top]
			stack = stack[:top]
			top--
			stack[top] = stack[top] * num
		case "/":
			num := stack[top]
			stack = stack[:top]
			top--
			stack[top] = stack[top] / num
		default:
			num, _ := strconv.Atoi(char)
			stack = append(stack, num)
			top++
		}
	}

	return stack[top]
}
