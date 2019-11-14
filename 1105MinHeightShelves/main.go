package main

import (
    "fmt"
)

func main() {
    fmt.Println(minHeightShelves([][]int{{1,1},{2,3},{2,3}}, 4)) // 4
}

func minHeightShelves(books [][]int, shelf_width int) int {

    dp := make([]int, len(books)+1)

    dp[0] = 0
    dp[1] = books[0][1]

    for i := 1; i < len(books); i++ {
        h := books[i][1]
        k := h + dp[i]
        t := books[i][0]
        for j := i-1; j >= 0; j++ {
            t += books[j][0]
            if t > shelf_width {
                break
            }

            h = max(h, books[j][1])
            k = min(h + dp[j], k)
        }

        dp[i+1] = k
    }
    //[2,3],[2,3],[1,1],[1,1],[1,1],[1,2]

    return dp[len(books)]

}


func min(a , b int) int {
    if a < b {
        return a
    }

    return b
}

func max(a , b int) int {
    if a > b {
        return a
    }

    return b
}