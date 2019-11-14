// FindDisappearedNumbers
package main

import "fmt"

func main() {
    fmt.Println(findDisappearedNumbers([]int{4,3,2,7,8,2,3,1}))
}

func findDisappearedNumbers(nums []int) []int {
    r := make([]int, 0)
    for i := 0; i < len(nums) ; i++ {
        v := nums[i] - 1
        for ; v >= 0; {
            t := nums[v] - 1
            nums[v] = 0
            if v == t {
                break
            }
            v = t
        }
    }

    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            r = append(r, i + 1)
        }
    }

    return r
}