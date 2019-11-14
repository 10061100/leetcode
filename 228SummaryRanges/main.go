package main

import (
    "fmt"
    "strconv"
)

func main() {
    fmt.Println(summaryRanges([]int{0,1,2,4,5,7}))
    fmt.Println(summaryRanges([]int{0,2,3,4,6,8,9}))
}

func summaryRanges(nums []int) []string {
    l := len(nums)

    res := make([]string, 0)
    if l <= 0 {
        return res
    }

    start, end := 0, 1
    for ; end <= l; end++{
        // 如果当前end已经和之前的杠不上了, 就汇总之前的
        if end == l || nums[end] - nums[end-1] != 1 {
            if end - start == 1 {
                res = append(res, strconv.Itoa(nums[start]))
            } else {
                res = append(res, strconv.Itoa(nums[start]) + "->" + strconv.Itoa(nums[end-1]))
            }
            start = end
        }
    }

    return res
}
