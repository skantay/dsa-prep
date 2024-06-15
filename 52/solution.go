func totalNQueens(n int) int {
	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
	}

	var result [][]string

	var backtrack func(int)

	backtrack = func(curRow int) {
		if curRow == n {
			var soln []string
			for r := 0; r < n; r++ {
				soln = append(soln, strings.Join(board[r], "."))
			}
			result = append(result, soln)
			return
		}

		for i := 0; i < n; i++ {
			if valid(board, curRow, i) {
				board[curRow][i] = "Q"

				backtrack(curRow + 1)

				board[curRow][i] = ""
			}
		}
	}

	backtrack(0)

	return len(result)
}

func valid(board [][]string, m, n int) bool {
	for v := m - 1; v >= 0; v-- {
		if board[v][n] == "Q" {
			return false
		}
	}

	for x, y := m-1, n-1; x >= 0 && y >= 0; x, y = x-1, y-1 {
		if board[x][y] == "Q" {
			return false
		}
	}

	for x, y := m-1, n+1; x >= 0 && y < len(board[0]); x, y = x-1, y+1 {
		if board[x][y] == "Q" {
			return false
		}
	}

	return true
}
