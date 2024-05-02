package leetcode

func maximumSumOfHeights(heights []int) int64 {    
    var maxSum int64 
    
    for i := 0; i < len(heights); i++ {
        maxSum = max(maxSum, getSum(i, heights))
    }

    return maxSum
}

func getSum(i int, heightsOriginal []int) int64 {
    heights := make([]int, len(heightsOriginal))
    copy(heights, heightsOriginal)

    stack := make([]int, 1)

    var rightSide int = i

    stack[0] = rightSide

    rightSide++

    for ; rightSide < len(heightsOriginal); rightSide++ {
        if len(stack) != 0 && heights[rightSide] > heights[stack[len(stack)-1]] {
            heights[rightSide] = heights[stack[len(stack)-1]]
        }

        stack = append(stack, rightSide)
    }

    stack = make([]int, 1)

    var leftSide int = i

    stack[0] = leftSide

    leftSide--

    for ; leftSide >= 0; leftSide-- {
        if len(stack) != 0 && heights[leftSide] > heights[stack[len(stack)-1]] {
            heights[leftSide] = heights[stack[len(stack)-1]]
        }

        stack = append(stack, leftSide)
    }

    var sum int64

    for ind := range heights {
        sum += int64(heights[ind])
    }

    return sum
}
