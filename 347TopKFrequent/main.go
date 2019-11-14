package main

import (
    "fmt"
    "sort"
)

func main() {
    fmt.Println(topKFrequent([]int{1,1,1,1,2}, 1))
}

func topKFrequent(nums []int, k int) []int {
    res := make([]int, 0)

    m := make(map[int]int)
    for _, v := range nums {
        if _, ok := m[v]; !ok {
            m[v] = 0
        }

        m[v]++
    }

    t := make([][]int, 0)
    for v, c := range m {
        t = append(t, []int{v, c})
    }

    sort.Sort(T(t))

    for i := 0; i < k && i < len(m); i++ {
        res = append(res, t[i][0])
    }

    return res
}


type T [][]int

func (t T) Len() int {
    return len(t)
}

func (t T) Less(i, j int) bool {
    ci := t[i][1]
    cj := t[j][1]

    if ci > cj {
        return true
    }

    return false
}

func (t T) Swap(i, j int) {
    t[i], t[j] = t[j], t[i]
}