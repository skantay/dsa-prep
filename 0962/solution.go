package leetcode

func maxWidthRamp(nums []int) int {
    startPoints := make([]int, 1, len(nums))

    for i := range nums {
        if nums[i] < nums[startPoints[len(startPoints)-1]] {
            startPoints = append(startPoints, i)
        } 
    }

    maxs := make([]int, 0) 

    for i := len(nums) - 1; i >= 0; i-- {
        for len(startPoints) != 0 && nums[i] >= nums[startPoints[len(startPoints)-1]] {
            maxs = append(maxs, i - startPoints[len(startPoints)-1])
            startPoints = startPoints[:len(startPoints)-1]
        }

        if len(startPoints) == 0 {
            break
        }
    }

    var max int

    for i := range maxs {
        if max < maxs[i] {
            max = maxs[i]
        }
    }

    return max
}
