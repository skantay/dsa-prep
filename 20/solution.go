package leetcode 

func isValid(s string) bool {
    if len(s) == 1 {
        return false
    }
    stack := make([]rune, 0)

    for i := 0; i < len(s);i++ {
        char := rune(s[i])
        switch char {
        case ')':
            if len(stack) != 0 && stack[len(stack)-1] == '(' {
                stack = stack[:len(stack)-1]
            } else {
                stack = append(stack, char)
            }
        case '}':
            if len(stack) != 0 && stack[len(stack)-1] == '{' {
                stack = stack[:len(stack)-1]
            } else {
                stack = append(stack, char)
            }
        case ']':
            if len(stack) != 0 && stack[len(stack)-1] == '[' {
                stack = stack[:len(stack)-1]
            } else {
                stack = append(stack, char)
            }
        default:
            stack = append(stack, char)
        }
    }

    return len(stack) == 0
}
