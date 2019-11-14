package main

import (
    "fmt"
)

func main()  {
    fmt.Println(largest1BorderedSquare([][]int{{1}}) == 1)
    fmt.Println(largest1BorderedSquare([][]int{{0,0}, {0,0}}) == 0)
    fmt.Println(largest1BorderedSquare([][]int{{1,1}, {1,1}}) == 4)
    fmt.Println(largest1BorderedSquare([][]int{{1,1, 1}, {1,1,1}, {1,1,1}}) == 9)
    fmt.Println(largest1BorderedSquare([][]int{{1,1, 1}, {1,0,1}, {1,1,1}}) == 9)
    fmt.Println(largest1BorderedSquare([][]int{{0,1, 1}, {1,1,1}, {1,1,1}}) == 4)
    fmt.Println(largest1BorderedSquare([][]int{{1,1,1,1}, {1,0,0,1}, {1,0,0,1}, {1,1,1,1}}) == 16)
    fmt.Println(largest1BorderedSquare([][]int{{0,1,1,1}, {1,1,1,1}, {1,1,1,1}, {1,1,1,1}}) == 9)
    fmt.Println(largest1BorderedSquare([][]int{{0,1,1,1}}) == 1)
}

func largest1BorderedSquare(grid [][]int) int {
    h := len(grid)
    w := len(grid[0])

    rows := make([][]int, h)
    cols := make([][]int, h)
    // dp   := make([][]int, h)

    res := 0
    for i := 0; i < h; i++ {
        rows[i] = make([]int, w)
        cols[i] = make([]int, w)
        // dp[i]   = make([]int ,w)

        rows[i][0] = grid[i][0]
        res = max(res, grid[i][0])
    }

    for i := 0; i < w; i++ {
        cols[0][i] = grid[0][i]
        res = max(res, grid[0][i])
    }

    // 先计算rows, 即每个元素, 向左边最长的连续1
    for i := 0; i < h; i++ {
        for j := 1; j < w; j++ {
            if grid[i][j] == 0 {
                rows[i][j] = 0
            } else {
                rows[i][j] = rows[i][j-1] + 1
            }
        }
    }

    // 再计算cols
    for i := 0; i < w; i++ {
        for j := 1; j < h; j++ {
            if grid[j][i] == 0 {
                cols[j][i] = 0
            } else {
                cols[j][i] = cols[j-1][i] + 1
            }
        }
    }


    // dp[0][0] = grid[0][0]
    for i := 1; i < h; i++ {
        for j := 1; j < w; j++ {
            if grid[i][j] != 1 {
                continue
            }
            res = max(res, 1)
            t := min(rows[i][j], cols[i][j])

            for k := t; k > 1; k-- {
                up  := i - k + 1
                left:= j - k + 1

                if rows[up][j] >= k && cols[i][left] >= k {
                    res = max(res, k)
                    break
                }
            }
        }
    }

    return res*res
}


func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}

func min(a, b int) int {
    if a > b {
        return b
    }

    return a
}
