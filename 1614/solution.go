package leetcode

func maxDepth(s string) int {
    stack := make([]rune, 0)

    var top int = -1

    var maxDepth int

    for _, char := range s {
        if char == '(' {
            stack = append(stack, char)
            top++
        } else if char == ')' {
            maxDepth = max(maxDepth, len(stack))
            stack = stack[:top]
            top--
        }
    }

    return maxDepth
}

