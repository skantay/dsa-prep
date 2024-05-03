package leetcode

func longestConsecutive(nums []int) int {
    slices.SortFunc(nums, func(a, b int) int {
        if a < b {
            return -1
        }

        return 1
    })

    var count, res int

    for i := range nums {
        if i == 0 || nums[i-1] == nums[i] - 1  {
            count++
        } else if nums[i-1] != nums[i] {
            res = max(res, count)
            count = 1
        }
    }

    res = max(res, count)

    return res
}
