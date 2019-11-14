package main


func maximalSquare(matrix [][]byte) int {
    y := len(matrix)
    if y == 0 {
        return 0
    }

    x := len(matrix[0])
    if x == 0 {
        return 0
    }

    dp := make([][]int, y+1)
    for i := 0; i <= y ; i++ {
        dp[i] = make([]int, x+1)
    }

    for i := 0; i < x+1; i++ {
        dp[0][i] = 0
    }

    for i := 0; i < y+1; i++ {
        dp[i][0] = 0
    }

    // 初始化dp
    for i := 0; i < y; i++ {
        for j := 0; j < x; j++ {

            t := 0
            if matrix[i][j] == 1 {
                t = 1
            }

            dp[i+1][j+1] = dp[i][j+1] + dp[i+1][j] - dp[i][j] + t
        }
    }

    res := 0
    for i := 0 ; i < y ; i++ {
        for j := 0; j < x ; j++ {
            if matrix[i][j] != 1 {
                continue
            }

            res = max(res, 1)
            for n := 1; i + 1 - n >= 0 && j + 1 - n >= 0; n++{
                t := (n+1) * (n+1)
                sum1 := dp[i+1][j+1] - dp[i+1-n][j+1] - dp[i+1][j+1-n] + dp[i+1-n][j+1-n]

                if sum1 != t {
                    break
                }

                res = max(res, t)
            }
        }
    }

    return res
}


func max(a , b int) int {
    if a > b {
        return a
    }

    return b
}