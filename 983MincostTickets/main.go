package main

import (
    "fmt"
)

func main()  {
    fmt.Println(mincostTickets([]int{1,4,6,7,8,20}, []int{2,7,15}) == 11)
    fmt.Println(mincostTickets([]int{1,2,3,4,5,6,7,8,9,10,30,31}, []int{2,7,15}) == 17)
    fmt.Println(mincostTickets([]int{1,2,3,4,6,8,9,10,13,14,16,17,19,21,24,26,27,28,29}, []int{3,14,50}) == 50)
}


func mincostTickets(days []int, costs []int) int {
    res := len(days) * costs[0]

    if res == 0 {
        return res
    }

    dp := make([]int, 366)
    m  := make([]int, 365)

    dp[days[0]] = min(min(costs[0], costs[1]), costs[2])
    for _, v := range days {
        m[v-1] = 1
    }

    // fmt.Println(m)
    for i := 0; i < days[0]; i++ {
        dp[i] = 0
    }

    for i := days[0]+1; i <= 365; i++ {
        dp[i] = -1
    }
    // fmt.Println(dp)
    subJob(dp, days[len(days)-1], m, costs)
    // fmt.Println(dp)
    return dp[days[len(days)-1]]
}


func subJob(dp []int, k int, m []int, cost []int) int {
    if (dp)[k] != -1 {
        return (dp)[k]
    }

    if m[k-1] == 0 {
        (dp)[k] = subJob(dp, k-1, m, cost)
        return (dp)[k]
    }

    t := subJob(dp, k-1, m, cost) + cost[0]
    if k >= 7 {
        t = min(t, subJob(dp, k-7, m, cost) + cost[1])
    } else {
        t = min(t, cost[1])
    }

    if k >= 30 {
        t = min(t, subJob(dp, k-30, m, cost) + cost[2])
    } else {
        t = min(t, cost[2])
    }

    (dp)[k] = t
    return (dp)[k]
}


func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}
