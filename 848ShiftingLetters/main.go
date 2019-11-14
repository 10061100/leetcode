package main

import (
    "fmt"
)

func main()  {
    fmt.Println(shiftingLetters("abc", []int{3,5,9}))
}

func shiftingLetters(S string, shifts []int) string {
    if len(S) == 0 {
        return S
    }

    dp := make([]int, len(shifts))
    dp[len(shifts) - 1] = shifts[len(shifts) - 1]

    for i := len(shifts) - 2; i >= 0; i-- {
        dp[i] = dp[i+1] + shifts[i]
    }

    res := make([]byte, len(S))

    for k, v := range S {
        res[k] = 'a' + byte((int(v - 'a') + dp[k]) % 26)
    }

    return string(res)
}
