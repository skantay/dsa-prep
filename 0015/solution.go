package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	result := make([][]int, 0)

	prevVal := nums[0]

	for i := 0; i < len(nums)-2; i++ {
		if i != 0 && nums[i] == prevVal {
			continue
		}
		j, k := i+1, len(nums)-1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]

			if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			} else {
				result = append(result, []int{nums[i], nums[j], nums[k]})

				for j < len(nums)-1 && nums[j] == nums[j+1] {
					j++
				}

				j++
			}
		}

		prevVal = nums[i]
	}

	return result
}
