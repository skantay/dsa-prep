package leetcode

import "sort"

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	var prevMin int

	for i := 0; i <= len(nums)-k; i++ {
		sl := nums[i : i+k]

		min := sl[len(sl)-1] - sl[0]

		if prevMin > min || i == 0 {
			prevMin = min
		}
	}

	return prevMin
}
