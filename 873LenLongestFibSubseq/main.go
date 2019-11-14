package main

import (
    "fmt"
)

func main() {
    fmt.Println(lenLongestFibSubseq([]int{1,2,3,4,5,6,7,8}) == 5)
    fmt.Println(lenLongestFibSubseq([]int{1,3,7,11,12,14,18}) == 3)
    fmt.Println(lenLongestFibSubseq([]int{1,3,7}) == -1)
}


func lenLongestFibSubseq(A []int) int {
    m := make(map[int]int)

    index := make(map[int]int)
    for i := 0; i < len(A); i++ {
        index[A[i]] = i
    }

    for i := 0; i < len(A); i++ {
        for j := 0; j < i ; j++ {
            m[j * len(A) + i] = 1
        }
    }

    res := 0
    for i := 0; i < len(A); i++ {
        for j := 0; j < i; j++ {
            t := A[i] - A[j]
            k, ok := index[t]
            // 找到了前置的数列
            if t < A[j] && ok {
                m[j * len(A) + i] = m[k *len(A) + j] + 1
                res = max(res, m[j*len(A) + i] + 1)
            }
        }
    }

    if res <= 2 {
        return  -1
    }

    return res
}


func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}