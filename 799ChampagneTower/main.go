package main

import (
    "fmt"
)

func main() {
    fmt.Println(champagneTower(100, 2, 2))
}


func champagneTower(poured int, query_row int, query_glass int) float64 {
    // 换个思路, 看看每个杯子流过的流量
    dp := make([][]float64, query_row+1)
    for i := 0; i < query_row+1; i++ {
        dp[i] = make([]float64, query_row+1)
    }

    dp[0][0] = float64(poured)

    for i := 1; i <= query_row; i++ {
        dp[i][0] = max((dp[i-1][0] - 1)/2, 0)
        dp[i][query_row] = max((dp[i-1][i-1] - 1)/2, 0)

        for j := 1; j < query_row; j++ {
            left := (dp[i-1][j-1] - 1)/2
            right:= (dp[i-1][j] - 1) / 2

            if left < 0 {
                left = 0
            }

            if right < 0 {
                right = 0
            }

            dp[i][j] = left + right
        }
    }

    fmt.Println(dp)
    if dp[query_row][query_glass] > 1 {
        return 1
    }

    return dp[query_row][query_glass]
}



func max(a, b float64) float64 {
    if a > b {
        return a
    }

    return b
}