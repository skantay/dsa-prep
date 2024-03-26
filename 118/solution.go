package leetcode

func generate(numRows int) [][]int {
	result := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		sl := make([]int, i+1)

		for j := 0; j < len(sl); j++ {
			if j == 0 || j == len(sl)-1 {
				sl[j] = 1

				continue
			}

			left, right := result[i-1][j-1], result[i-1][j]

			sl[j] = left + right
		}

		result[i] = sl
	}

	return result
}
