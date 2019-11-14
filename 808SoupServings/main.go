package main

import (
    "fmt"
)

func main() {
    fmt.Println(soupServings(1000))
}


func soupServings(N int) float64 {

    // 因为只有每次都是 25, 75的情况下, 才会是B先完
    if N >= 5001 {
        return 1.0
    }


    // 这个时候, 用存储中间结果来进行
    n := (N+24)/25
    dp := make([][][]float64, 3) // 分别是 0 = A先完成, 1 = AB一起完成, 2 = B先完成

    // 初始化dp数组
    for i := 0; i < len(dp); i++ {
        dp[i] = make([][]float64, n+1)
        for j := 0; j < n + 1; j++ {
            dp[i][j] = make([]float64, n+1)
        }
    }

    dp[1][0][0] = 1 // 当没有的时候, AB一起完成的概率是1, 其余都是0

    // 如果 A 是0, B 不是0, 则一定是A先完成
    for i := 1; i <= n; i++ {
        dp[0][0][i] = 1
    }

    t := [][]int{{4, 0}, {3, 1}, {2, 2}, {1, 3}}
    for i := 1; i <= n; i++ {
        for j := 1; j <= n; j++ {

            // p1 := float64(0) // A先完成
            // p2 := float64(0) // AB一起完成
            for _, v := range t {
                newx := i - v[0]
                newy := j - v[1]

                if newx <= 0 {
                    newx = 0
                }

                if newy <= 0 {
                    newy = 0
                }
                dp[0][i][j] += float64(0.25) * dp[0][newx][newy]
                dp[1][i][j] += float64(0.25) * dp[1][newx][newy]
            }
        }
    }

    return dp[0][n][n] + dp[1][n][n] / 2

}
