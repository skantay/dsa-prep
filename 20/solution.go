package leetcode 

func isValid(s string) bool {
    if len(s) == 1 {
        return false
    }

    var top int = -1

    stack := make([]rune, 0)

    for i := 0; i < len(s);i++ {
        char := rune(s[i])

        switch char {
        case ')':
            if len(stack) != 0 && stack[top] == '(' {
                stack = stack[:top]
                top--
            } else {
                stack = append(stack, char)
                top++
            }
        case '}':
            if len(stack) != 0 && stack[top] == '{' {
                stack = stack[:top]
                top--
            } else {
                stack = append(stack, char)
                top++
            }
        case ']':
            if len(stack) != 0 && stack[top] == '[' {
                stack = stack[:top]
                top--
            } else {
                stack = append(stack, char)
                top++
            }
        default:
            stack = append(stack, char)
            top++
        }
    }

    return len(stack) == 0
}
