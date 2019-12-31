package main

import (
    "fmt"
)

func main()  {

    fmt.Println(mctFromLeafValues([]int{6,2,4,8,1,5}) == 125)
    fmt.Println(mctFromLeafValues([]int{2,4,8,1,5}) == 85)
    fmt.Println(mctFromLeafValues([]int{4,8,1,5}) == 77)
}


func mctFromLeafValues(arr []int) int {
    if len(arr) <= 2 {
        return arr[0]*arr[1]
    }

    n := len(arr)

    dp := make([][]int, n)
    m  := make([][]int, n)

    for i := 0; i < n; i++ {
        dp[i] = make([]int, n)
        m[i] = make([]int, n)
    }

    copy(m[0], arr)

    for i:= 1; i < n; i++ {
        for j := 0; j + i < n; j++ {
            dp[i][j] = m[0][j] * m[i-1][j+1] + dp[i-1][j+1]
            for k := 1; k < i; k++ {
                t := dp[k][j] + dp[i-k-1][j+k+1] + m[k][j] * m[i-k-1][j+k+1]
                dp[i][j] = min(dp[i][j] , t)
            }
            // right:= m[i-1][j] * m[0][j+i] + dp[i-1][j]

        }
        for j := 0; j + i < n; j++ {
            m[i][j] = max(m[i-1][j], m[0][j+i])
        }
    }

    // fmt.Println(m)
    // fmt.Println(dp)

    return dp[n-1][0]
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}

func min(a, b int) int {
    if a > b {
        return b
    }

    return a
}