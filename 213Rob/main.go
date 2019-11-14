package main

import (
    "fmt"
)

func main() {
    fmt.Println(rob([]int{2,3,2}) == 3)
    fmt.Println(rob([]int{1,2,3,1}) == 4)
    fmt.Println(rob([]int{1,1}) == 1)
    fmt.Println(rob([]int{1,2}) == 2)
}


func rob(nums []int) int {
    l := len(nums)

    if l == 0 {
        return 0
    }

    if l == 1 {
        return nums[0]
    }

    rfirst := []int{0, nums[0]}
    nfirst := []int{0, 0}

    for i := 1; i < l-1; i++ {
        // 抢劫第一家
        n := rfirst[0]
        r := rfirst[1]

        rfirst[0] = max(r, n) // 不抢当前人家
        rfirst[1] = n + nums[i]  // 抢了当前人家

        // 没抢第一家
        n = nfirst[0]
        r = nfirst[1]

        nfirst[0] = max(r, n) // 不抢当前人家
        nfirst[1] = n + nums[i]  // 抢了当前人家
    }

    // 抢劫了第一家, 最后一家就不能抢了
    r := max(rfirst[0], rfirst[1])

    // 没抢劫第一家, 最后一家还可以抢
    r = max(r, max(nfirst[0] + nums[l-1], nfirst[1]))

    return r
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
