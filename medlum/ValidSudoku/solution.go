func isValidSudoku(board [][]byte) bool {
    for i:=0;i<9;i++ {
        row := make([]int,9)
        col := make([]int,9)
        cube:= make([]int,9)
        for j:=0;j<9;j++ {
            if board[i][j] != '.' {
                if row[board[i][j] - '1'] == 1 {
                    return false
                } else {
                    row[board[i][j] - '1'] = 1
                }
            }
            if board[j][i] != '.' {
                if col[board[j][i] - '1'] == 1 {
                    return false
                } else {
                    col[board[j][i] - '1'] = 1
                }
            }
            cubeX := 3 * (i / 3) + j / 3
            cubeY := 3 * (i % 3) + j % 3
            if board[cubeX][cubeY] != '.' {
                if cube[board[cubeX][cubeY] - '1'] == 1 {
                    return false
                } else {
                    cube[board[cubeX][cubeY] - '1'] = 1
                }
            }
        }
    }
    return true
}