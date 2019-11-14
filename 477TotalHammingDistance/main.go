package main

import (
    "fmt"
)

func main() {
    fmt.Println(totalHammingDistance([]int{4,14,2}))
}

func totalHammingDistance(nums []int) int {
    res := 0

    if len(nums) <= 1 {
        return res
    }

    for ; ;  {
        isbreak := true
        n1 := 0
        n0 := 0
        for i := 0; i < len(nums); i++ {
            if nums[i] % 2 == 1 {
                n1++
            } else {
                n0++
            }

            nums[i] = nums[i] >> 1
            if nums[i] != 0 {
                isbreak = false
            }
        }

        res += n1 * n0

        if isbreak {
            break
        }

    }

    return res
}