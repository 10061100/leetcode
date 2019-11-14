package main

import (
    "fmt"
    "sort"
)

func main()  {
    fmt.Println(largestValsFromLabels([]int{5,4,3,2,1}, []int{1,3,3,3,2}, 3, 2))
    fmt.Println(largestValsFromLabels([]int{9,8,8,7,6}, []int{0,0,0,1,1}, 3, 1))
}

func largestValsFromLabels(values []int, labels []int, num_wanted int, use_limit int) int {
    t := make([][]int, 0)

    for k ,v := range values {
        t = append(t, []int{v, labels[k]})
    }

    sort.Slice(t, func(i, j int) bool {
        if t[i][0] > t[j][0] {
            return true
        }

        return false
    })

    v := 0
    m := make(map[int]int)
    for i := 0; i < len(t) && num_wanted > 0; i++ {
        // 先判断是否超过了use_limit
        if vc, ok := m[t[i][1]]; !ok || vc < use_limit {
            v += t[i][0]
            num_wanted--
            if !ok {
                m[t[i][1]] = 0
            }
            m[t[i][1]]++
        }
    }

    return v
}
