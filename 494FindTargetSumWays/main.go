package main

import (
    "fmt"
)

func main()  {
    println(findTargetSumWays2([]int{1, 1, 1, 1, 1}, 3))
    println(findTargetSumWays2([]int{1, 1, 1, 1, 1}, 5))
    println(findTargetSumWays2([]int{1, 0}, 1))
}


func findTargetSumWays2(nums []int, S int) int {
    l := len(nums)
    if l == 0 {
        return 0
    }

    dp := make([]map[int]int, l)
    for i := 0; i < l; i++ {
        dp[i]= make(map[int]int)
    }
    dp[0][nums[0]] = 1
    dp[0][-nums[0]] += 1

    for length := 1; length <= l-1; length++ {
        left := dp[length-1]
        for value, count := range left {
            tmp := []int{nums[length], -nums[length]}
            for _, v := range tmp {
                newvalue := value + v
                if _, ok := dp[length][newvalue]; !ok {
                    dp[length][newvalue] = 0
                }
                dp[length][newvalue] += count
            }
        }
    }

    // fmt.Println(dp[0][l-1])
    if v, ok := dp[l-1][S]; ok {
        return v
    }

    return 0
}


func findTargetSumWays(nums []int, S int) int {
    l := len(nums)
    if l == 0 {
        return 0
    }

    dp := make([][]map[int]int, l)
    for i := 0; i < l; i++ {
        dp[i] = make([]map[int]int, l)
        for j := 0; j < l; j++ {
            dp[i][j] = make(map[int]int)
        }

        dp[i][i][nums[i]] = 1
        dp[i][i][-nums[i]]= 1
    }

    fmt.Println(dp)

    for length := 1; length <= l-1; length++ {
        for start := 0; start + length <= l-1; start++ {
            end := start + length
            left := dp[start][end-1]
            right:= dp[end][end]
            for leftvalue, leftcount := range left {
                for rightvalue, rightcount := range right {
                    newvalue := leftvalue + rightvalue
                    newcount := leftcount * rightcount
                    // println(newvalue, ",", newcount)
                    if _, ok := dp[start][end][newvalue]; !ok {
                        dp[start][end][newvalue] = 0
                    }
                    dp[start][end][newvalue] += newcount
                }
            }
        }
    }

    fmt.Println(dp)
    // fmt.Println(dp[0][l-1])

    if v, ok := dp[0][l-1][S]; ok {
        return v
    }

    return 0
}
