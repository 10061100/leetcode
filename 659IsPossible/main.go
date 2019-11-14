package main

import (
    "fmt"
)

func main() {
    fmt.Println(isPossible([]int{1,2,3,3,4,5}) == true)
    fmt.Println(isPossible([]int{1,2,3,4,5}) == true)
    fmt.Println(isPossible([]int{1,2,3,4,4,5}) == false)
    fmt.Println(isPossible([]int{1,2,3,4,4,5,5,6}) == true)
}


func isPossible(nums []int) bool {

    l := len(nums)

    if l <= 2 {
        return false
    }

    m := make(map[int]int)

    // 预处理数据
    for i := 0; i < l; i++ {
        if _, ok := m[nums[i]]; !ok {
            m[nums[i]] = 0
        }

        m[nums[i]]++
    }

    end := make(map[int]int)
    for k := range m {
        end[k] = 0
    }


    for i := 0; i < l; i++ {
        t := nums[i]
        if m[t] <= 0 {
            continue
        }

        m[t]--

        // 如果自己可以和前一个数形成递增关系, 则关系链扩展
        if end[t-1] > 0 {
            end[t-1]--
            end[t]++
            continue
        }

        if m[t+1] <= 0 || m[t+2] <= 0 {
            // fmt.Println(m, ",", t)
            return false
        }

        m[t+1]--
        m[t+2]--
        end[t+2]++
    }

    return true
}


