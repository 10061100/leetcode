package main

import (
    "fmt"
)

func main() {
    fmt.Println(findMaxLength([]int{0,1,0,1}))
}


// 事实证明, 窗口方法是无效的, 类似于
// 1100000000011 这种case
func findMaxLength(nums []int) int {

    sum := 0
    m := make(map[int]int)
    m[0] = -1
    max := 0

    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {
            sum--
        } else {
            sum++
        }

        if v, ok := m[sum]; !ok {
            m[sum] = i
        } else {
            if t := i - v; t > max {
                max = t
            }
        }

    }

    return max
}
