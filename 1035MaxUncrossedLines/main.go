package main

import (
    "fmt"
)

func main()  {

    fmt.Println(maxUncrossedLines([]int{1,4,2}, []int{1,2,4}) == 2)
    fmt.Println(maxUncrossedLines([]int{2,5,1,2,5}, []int{10,5,2,1,5,2}) == 3)
    fmt.Println(maxUncrossedLines([]int{1,3,7,1,7,5}, []int{1,9,2,5,1}) == 2)
}


func maxUncrossedLines(A []int, B []int) int {
    la := len(A)
    lb := len(B)

    if la == 0 || lb == 0 {
        return 0
    }

    dp := make([][]int, la+1)
    for i := 0; i <= la; i++ {
        dp[i] = make([]int, lb+1)
    }

    for i := 0; i <= la; i++ {
        dp[i][0] = 0
    }

    for i := 0; i <= lb; i++ {
        dp[0][lb] = 0
    }

    for i := 0; i < la; i++ {
        a := A[i]
        for j := 0; j < lb; j++ {
            t1 := dp[i+1][j+1]
            t2 := dp[i][j+1]
            t3 := dp[i+1][j]
            dp[i+1][j+1] = max(t1, max(t2, t3))
            if a == B[j] {
                dp[i+1][j+1] = max(dp[i+1][j+1], dp[i][j] + 1)
            }
        }
    }

    return dp[la][lb]
}


func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}