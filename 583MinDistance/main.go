package main

import (
    "fmt"
)

func main() {
    fmt.Println(minDistance("ab", "abc"))
}

func minDistance(word1 string, word2 string) int {

    l1 := len(word1)
    l2 := len(word2)

    dp := make([]int, l2+1)
    tmp:= make([]int, l2+1)

    dp[0] = 0
    for i := 0; i < l2; i++ {
        dp[i+1] = i+1
    }
    copy(tmp, dp)

    for i := 0; i < l1; i++ {
        dp[0] = i+1
        for j := 0; j < l2; j++ {
            // 删除j && 删除i
            t := min(dp[j+1]+1, dp[j] + 1)
            // 两个相等, 等于i-1, j-1
            if word1[i] == word2[j] {
                t = min(tmp[j], t)
            }

            dp[j+1] = t
        }

        copy(tmp, dp)
    }

    return dp[l2]
}

func min(a, b int) int {
    if a > b {
        return b
    }

    return a
}
