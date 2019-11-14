package main

import (
    "fmt"
)

func main() {
    fmt.Println(lastStoneWeightII([]int{1,1,1,7}) == 4)
}


func lastStoneWeightII(stones []int) int {

    sum := 0
    for i := 0; i < len(stones); i++ {
        sum += stones[i]
    }

    // 动态规划, 保证两边的和最近
    c := (sum+1)/2 // 保证是大的值
    dp := make([]int, c+1)
    for i := 0; i < len(stones); i++ {
        for j := c; j >= stones[i]; j-- {
            dp[j] = max(dp[j], dp[j-stones[i]] + stones[i])
        }
    }

    return sum - 2 * dp[c]
}


func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}


