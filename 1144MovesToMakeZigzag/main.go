package main

import (
    "fmt"
)

func main()  {

    fmt.Println(movesToMakeZigzag([]int{9,6,1,6,2}) == 4)
    fmt.Println(movesToMakeZigzag([]int{1,2,3}) == 2)
}

func movesToMakeZigzag(nums []int) int {
    res1 := 0
    res2 := 0
    n := len(nums)

    if len(nums) == 1 {
        return 0
    }

    for i := 0; i < len(nums); i++ {
        // 说明偶数小于奇数
        if i % 2 == 0  {
            // 偶数位
            res1 += diff(nums, i, n)
        } else {
            res2 += diff(nums, i, n)
        }

    }

    return min(res1, res2)
}

func diff(nums []int, cur , n int) int {
    left := 0
    if cur < n-1 && nums[cur+1] <= nums[cur] {
        left = nums[cur] - nums[cur+1] + 1
        // nums[cur] = nums[cur+1] - 1
    }

    right := 0
    if cur >= 1 && nums[cur-1] <= nums[cur]{
        right = nums[cur] - nums[cur-1] + 1
        // nums[cur] = nums[cur-1] - 1
    }

    return max(left, right)
}


func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}


func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}

