package leetcode

func isValidSudoku(board [][]byte) bool {
    m := make(map[byte][][]int)

    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if board[i][j] == '.' {
                continue
            }

            locs := m[board[i][j]]

            for k := range locs {
                if locs[k][0] == i {
                    return false
                }
                
                if locs[k][1] == j {
                    return false
                }
                
                if locs[k][0]/3 == i/3 && locs[k][1]/3 == j/3 {
					return false
				}
            }

			m[board[i][j]] = append(m[board[i][j]], []int{i, j})
        }
    }

    return true
}
