package main

import (
    "fmt"
)

func main() {
    // fmt.Println(mergeStones([]int{3,5,1,2,6}, 3))
    // fmt.Println(mergeStones([]int{3,5,1,2,6}, 3))
    fmt.Println(mergeStones([]int{3,5,1,2,6}, 3))
}

func mergeStones(stones []int, K int) int {
    l := len(stones)

    if (l-1) % (K-1) != 0 {
        return -1
    }

    dp := make([][]int, l)
    sum:= make([][]int, l)
    for i := 0; i < l; i++ {
        dp[i] = make([]int, l)
        dp[i][i] = stones[i]
        sum[i] = make([]int, l)
        // sum[i][i] = stones[i]
    }

    for n := 1; ; n++ {
        length := n * (K-1) + 1
        if length > l {
            break
        }

        for start := 0; start + length - 1 < l; start++ {
            end := start + length - 1
            calSubOrder(K, K, dp, sum, start, end)
            fmt.Println(sum)
        }

    }

    return sum[0][l-1]
}


func calSubOrder(K int, cur int, dp [][]int, sum [][]int,  start, end int) (int, int) {

    if dp[start][end] != 0 {
        return dp[start][end], sum[start][end]
    }

    l := end - start + 1
    // 否则就去计算当下的值
    min := 0
    smin := 0
    for seg := 1;  l - seg >= cur - 1; seg += (K-1) {
        tv := dp[start][start+seg-1]
        after , cost := 0, 0
        if cur > 0 {
            after , cost = calSubOrder(K, cur-1, dp, sum, start+seg, end)
            // fmt.Println(after, ",", cost)
        }
        if min <= 0 || tv + after < min {
            min = tv + after
            smin = cost
        }

    }

    // fmt.Println(smin)

    sum[start][end] = smin
    dp[start][end] = min
    return min, smin
}
