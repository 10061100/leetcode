package main

import (
    "fmt"
)

func main() {
    fmt.Println(minFlipsMonoIncr("00110") == 1)
    fmt.Println(minFlipsMonoIncr("010110") == 2)
    fmt.Println(minFlipsMonoIncr("00011000") == 2)
}

func minFlipsMonoIncr(S string) int {
    l := len(S)

    if l <= 1 {
        return 0
    }

    // 每个字符右边全变1需要变化的次数(包括自己)
    dp := make([]int, l+1)
    dp[l] = 0
    for j := l-1; j >= 0; j-- {
        if S[j] == '0' {
            dp[j] = dp[j+1] + 1
        } else {
            dp[j] = dp[j+1]
        }
    }

    res := dp[0]
    pre := 0

    // 每个数, 左边全是0的变化值, 包括自己
    for i := 0; i < l; i++ {
        if S[i] == '1' {
            pre++
        }

        t := pre + dp[i+1]

        if t < res {
            res = t
        }
    }

    return res
}