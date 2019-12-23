package main

import (
    "fmt"
)

func main() {
    fmt.Println(minScoreTriangulation([]int{1,2,3}) == 6)
    fmt.Println(minScoreTriangulation([]int{1,3,1,4,1,5}) == 13)
}

func minScoreTriangulation(A []int) int {

    n := len(A)

    if n == 3 {
        return A[0] * A[1] * A[2]
    }

    cache := make([]int, n*n+1)

    return subOrder(cache, A, 0, n-1, n)

}

func subOrder(cache []int, A []int, a, b int, n int) int {

    if b == a + 1 {
        return 0
    }

    t := (a)*n + b

    if cache[t] != 0 {
        return cache[t]
    }

    m := 100000000
    for i := a + 1; i < b; i++ {
        m = min(m, A[a]*A[b]*A[i] + subOrder(cache, A, a, i, n) + subOrder(cache, A, i, b, n))
    }

    cache[t] = m

    return m
}

func min(a, b int) int {
    if a > b {
        return b
    }

    return a
}

