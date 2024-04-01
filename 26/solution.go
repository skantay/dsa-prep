package leetcode

func removeDuplicates(nums []int) int {
	i, j := 0, 0

	for j != len(nums) {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}

		j++
	}

	return i + 1
}
