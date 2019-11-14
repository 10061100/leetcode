package main

import (
    "fmt"
)

func main() {
    fmt.Println(maxTurbulenceSize([]int{9,4,2,10,7,8,8,1,9}))
    fmt.Println(maxTurbulenceSize([]int{1,2,1,3}))
    fmt.Println(maxTurbulenceSize([]int{1,1,1,1}))
    fmt.Println(maxTurbulenceSize([]int{1,2,3,4}))
    fmt.Println(maxTurbulenceSize([]int{9,9}))
}


func maxTurbulenceSize(A []int) int {
    if len(A) <= 2 {
        return len(A)
    }

    dp := make([]int, len(A)-1)
    for i := 0; i < len(A)-1; i++ {
        dp[i] = 1
        if A[i] < A[i+1] {
            dp[i] = -1
        } else if A[i] == A[i+1] {
            dp[i] = 0
        }
    }

    // fmt.Println(dp)
    start, end := 0, 0
    res := 0
    for ; end < len(dp); end++ {

        if dp[end] == 0 {
            start = end+1
            continue
        }

        if start == end {
            continue
        }

        if dp[end] != dp[end-1] {
            res = max(res, end - start + 1)
        } else {
            res = max(1, res)
            start = end
        }
    }

    return res + 1
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
