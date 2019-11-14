package main

import (
    "fmt"
)

func main()  {
    // fmt.Println(largestSumOfAverages([]int{0,1,2}, 3))
    fmt.Println(largestSumOfAverages([]int{9,1,2,3, 9}, 3))
}


func largestSumOfAverages(A []int, K int) float64 {
    l := len(A)
    sum := make([]int, l+1)
    sum[0] = 0
    for i := 1; i <= l; i++ {
        sum[i] = sum[i-1] + A[i-1]
    }

    if l <= K {
        return float64(sum[l])
    }

    if K == 1 {
        return float64(sum[l])/float64(l)
    }

    dp := make([]float64, l)
    tmp:= make([]float64, l)

    for i := 0; i < l; i++ {
        dp[i] = float64(sum[i+1])/float64(i+1)
    }

    for i := 1; i < K ; i++ {
        for j := i; j < l; j++ {
            tmp[j] = dp[j-1] + float64(A[j])
            for t := j-1; t >= i-1 ; t-- {
                tmp[j] = max(tmp[j], dp[t] + float64(sum[j+1] - sum[t+1])/float64(j-t))
            }
        }

        copy(dp, tmp)
    }

    return dp[l-1]
}


func max(a, b float64) float64 {
    if a > b {
        return a
    }

    return b
}