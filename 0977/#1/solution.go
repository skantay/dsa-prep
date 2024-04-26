package leetcode

func sortedSquares(nums []int) []int {
	for i, n := range nums {
		nums[i] = n * n
	}
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	return nums
}
