package main

import (
    "fmt"
)

func main() {
    fmt.Println(totalFruit([]int{0,1,2,2}))
    fmt.Println(totalFruit([]int{3,3,3,1,2,1,1,2,3,3,4}))
    fmt.Println(totalFruit([]int{1,2}))
    fmt.Println(totalFruit([]int{1,2,3,2,2}))
}

func totalFruit(tree []int) int {

    m := make(map[int]int)
    left := 0
    right:= 0
    m[tree[0]] = 0

    res := 0
    for ; right < len(tree); {
        // 首先right往后面走, 直到m大小超过三
        for ; right < len(tree); right++ {
            if _, ok := m[tree[right]]; !ok {
                m[tree[right]] = 0
            }
            m[tree[right]]++

            // 做多两种
            if len(m) <= 2 {
                res = max(res, right - left + 1)
                continue
            }

            // 超过两种了, 左指针一致往前走, 直到len(m)<=2
            for ; left < right && len(m) >= 3; left++{
                m[tree[left]]--
                if m[tree[left]] <= 0 {
                    delete(m, tree[left])
                    // fmt.Println(m)
                    // fmt.Println(left)
                }
            }

            // fmt.Println(left, ",", right)

        }

        // fmt.Println(left, ",", right)
    }

    return res
}


func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}