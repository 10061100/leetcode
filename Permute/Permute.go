// Permute
package main

import (
    "fmt"
)

func main() {
    a := []int{3,3,0,3}
    sort(&a)
    fmt.Println(a)
}

func permute(nums []int) [][]int {
    r := make([][]int, 0)
    k := 1
    s := make([]int, len(nums))

    permuteSingle(&r, nums, k, s)

    return r
}


func permuteSingle(r *[][]int, nums []int, k int, s []int) {
    if k == len(nums) {
        s[k-1] = nums[k-1]
        c := make([]int, k)
        copy(c, s)
        *r = append(*r, c)
        return
    }

    ex := make(map[int]int)

    for i := k - 1; i < len(nums); i++ {

        // 跳过所有和自己想通的数
        if i != k - 1 {
            if _, ok := ex[nums[i]]; ok {
                // 说明和当前的
                continue
            }
        }

        ex[nums[i]] = 1

        if i != k - 1 {
            t := nums[i]
            nums[i] = nums[k-1]
            nums[k-1] = t
        }

        s[k-1] = nums[k-1]
        permuteSingle(r, nums, k+1, s)

        // 调换回来
        if i != k - 1 {
            t := nums[i]
            nums[i] = nums[k-1]
            nums[k-1] = t
        }
    }
}

func sort(nums *[]int) {
    for i := 0 ; i < len(*nums) ; i++ {
        for j := 0; j < len(*nums) - 1; j++ {
            if (*nums)[j+1] < (*nums)[j] {
                t := (*nums)[j]
                (*nums)[j] = (*nums)[j+1]
                (*nums)[j+1] = t
            }
        }
    }
}