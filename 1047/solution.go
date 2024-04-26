package leetcode

func removeDuplicates(s string) string {
    stack := make([]rune, 0)

    for _, char := range s {
        if len(stack) != 0 && char == stack[len(stack)-1] {
            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, char)
        }
    }

    return string(stack)
}
