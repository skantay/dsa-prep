package leetcode

func findMaxConsecutiveOnes(nums []int) int {
    var max int
    var count int
    for i := 0; i < len(nums); i++ {
        if nums[i] == 1 {
            count++
        } else {
            count = 0
        }
        if max < count  {
            max = count
        }
    }

    return max
}