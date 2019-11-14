package main

import (
    "fmt"
    "sort"
)

func main() {
    fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}

func fourSum(nums []int, target int) [][]int {
    sort.Sort(sort.IntSlice(nums))
    return NSum(nums, target, 4, 0)
}


func NSum(nums []int, target int, N int, start int) [][]int {
    res := make([][]int, 0)

    if start >= len(nums) {
        return res
    }

    // 滑动窗口
    if N == 2 {
        for l, r := start, len(nums) - 1; l < r ; {
            if nums[l] + nums[r] == target {
                res = append(res, []int{nums[l], nums[r]})
                for l++ ; l < len(nums) && nums[l] == nums[l-1]; l++ {} // 跳过自己相同的
            }

            if nums[l] + nums[r] < target {
                l++
            } else {
                r--
            }
        }

    } else {
        // 分解后递归
        for l := start; l < len(nums); {
            nt := target - nums[l]
            t := NSum(nums, nt, N-1, l+1)
            for _, v := range t {
                res = append(res, append([]int{nums[l]}, v...))
            }
            for l++ ; l < len(nums) && nums[l] == nums[l-1]; l++ {} // 跳过自己相同的
        }
    }

    return res
}
