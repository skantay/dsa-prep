package leetcode

func luckyNumbers(matrix [][]int) []int {
	var res []int
	var candidates []int
	for i := range matrix {
		candID := 0
		for j := range matrix[i] {
			if matrix[i][j] < matrix[i][candID] {
				candID = j
			}
		}
		candidates = append(candidates, candID)
	}
	for i := range candidates {
		isReal := true
		for j := range matrix {
			if matrix[j][candidates[i]] > matrix[i][candidates[i]] {
				isReal = false
				break
			}
		}
		if isReal {
			res = append(res, matrix[i][candidates[i]])
		}
	}
	return res
}
