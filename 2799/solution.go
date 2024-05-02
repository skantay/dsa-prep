package leetcode

func countCompleteSubarrays(nums []int) int {
	distinct := make(map[int]struct{})

	for i := range nums {
		distinct[nums[i]] = struct{}{}
	}

	var countDistinct int = len(distinct)

	var l, r, subArrays int

	counter := make(map[int]int)

	for r < len(nums) {
		counter[nums[r]]++

		for l <= r && countDistinct == len(counter) {
			counter[nums[l]]--

			if counter[nums[l]] == 0 {
				delete(counter, nums[l])
			}

			subArrays += len(nums) - r
			l++
		}

		r++
	}

	return subArrays
}
