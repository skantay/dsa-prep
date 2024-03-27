package leetcode

func maxArea(height []int) int {
	i, j := 0, len(height)-1

	var prevMax int

	for i < j {
		min := height[i]

		if min > height[j] {
			min = height[j]
		}

		s := min * (j - i)

		if s > prevMax {
			prevMax = s
		}

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return prevMax
}
