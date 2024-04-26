package leetcode

func validateStackSequences(pushed []int, popped []int) bool {

    stack := make([]int, 0, len(pushed))

    var pushedInd, poppedInd int

    for i := 0; i < len(pushed) * 2; i++ {
        if len(stack) != 0 && stack[len(stack)-1] == popped[poppedInd] {
            poppedInd++
            stack = stack[:len(stack)-1]
            i = 0
        } else if pushedInd < len(pushed) {
            stack = append(stack, pushed[pushedInd])
            pushedInd++
        }
    }


    return len(stack) == 0 
}
