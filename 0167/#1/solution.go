package leetcode

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if target == numbers[i]+numbers[j] {
				return []int{i + 1, j + 1}
			}
		}
	}

	return nil
}
