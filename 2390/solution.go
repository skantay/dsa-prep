package leetcode

func removeStars(s string) string {
    stack := make([]rune, 0, len(s))

    for _, char := range s {
        if char == '*' {
            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, char)
        }
    }

    return string(stack)
}
