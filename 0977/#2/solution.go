package leetcode

func sortedSquares(nums []int) []int {
	i, j := 0, len(nums)-1
	result := make([]int, len(nums))
	ind := len(nums) - 1

	for i <= j {
		if abs(nums[i]) > abs(nums[j]) {
			result[ind] = nums[i] * nums[i]
			i++
		} else {
			result[ind] = nums[j] * nums[j]
			j--
		}

		ind--
	}

	return result
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}

	return n
}
