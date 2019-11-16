package main

import (
    "fmt"
)

func main()  {

    fmt.Println(maxSumAfterPartitioning([]int{1,15,7,9,2,5,10}, 3))
}

// 不多说了, 基础dp
func maxSumAfterPartitioning(A []int, K int) int {
    n := len(A)
    dp := make([]int, n+1)

    dp[0] = 0
    dp[1] = A[0]

    for i := 1; i < n; i++ {
        t := A[i]
        sum := t + dp[i]
        for j := 1; j < K && i - j >= 0; j++ {
            t = max(t, A[i-j])
            sum = max(sum , t * (j+1) + dp[i-j])
        }
        dp[i+1] = sum
    }

    return dp[n]
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
