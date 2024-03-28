package leetcode

func sortedSquares(nums []int) []int {
	for i, n := range nums {
		nums[i] = n * n
	}

	for i := 0; i < len(nums); i++ {

		minInd := i

		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minInd] {
				minInd = j
			}
		}

		nums[i], nums[minInd] = nums[minInd], nums[i]
	}

	return nums
}
