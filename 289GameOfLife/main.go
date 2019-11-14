package main

func gameOfLife(board [][]int)  {

    t := [][]int{{1,-1},{1,0},{1,1},{0,1},{-1,1},{-1,0},{-1,-1},{0,-1}}

    m := len(board)
    if m <= 0 {
        return
    }

    n := len(board[0])
    if n <= 0 {
        return
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            liveCount := 0
            deadCount := 0
            for _, v := range t {
                x := i + v[0]
                y := j + v[1]

                if x < 0 || x >= m {
                    continue
                }

                if y < 0 || y >= n {
                    continue
                }

                // 0 是死细胞, 1 是活细胞, 2 0->1, 由死变活, 3 1->0 由活变死
                if board[x][y] == 0 || board[x][y] == 2 {
                    deadCount++
                } else if board[x][y] == 1 || board[x][y] == 3 {
                    liveCount++
                }
            }

            if board[i][j] == 0 {
                // 当前是死细胞
                if liveCount == 3 {
                    board[i][j] = 2
                }
            } else {
                // 当前是活细胞
                if liveCount < 2 {
                    board[i][j] = 3
                } else if liveCount <= 3 {
                    continue
                } else {
                    board[i][j] = 3
                }
            }

        }
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == 2 {
                board[i][j] = 1
            } else if board[i][j] == 3 {
                board[i][j] = 0
            }
        }
    }
}
