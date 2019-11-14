package main

import (
    "fmt"
)

func main() {
    fmt.Println(fourSumCount([]int{1,2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}))
}

func fourSumCount(A []int, B []int, C []int, D []int) int {
    tmp := make(map[int]int)

    for _, a := range A {
        for _, b := range B {
            if _, ok := tmp[a+b]; !ok {
                tmp[a+b] = 0
            }

            tmp[a+b]++
        }
    }

    res := 0
    for _, c := range C {
        for _, d := range D {
            t := 0 - (c+d)
            if v, ok := tmp[t]; ok {
                res += v
            }
        }
    }

    return res
}
