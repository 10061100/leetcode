package main

import (
    "fmt"
)

func main() {
    fmt.Println(check(constructArray(6,5)))
    fmt.Println(check(constructArray(5,1)))
}


func check(a []int) int {
    m := make(map[int]int)
    for i := 0; i + 1 < len(a); i++ {
        t := a[i+1] - a[i]
        if t < 0 {
            t = -t
        }
        m[t]++
    }

    return len(m)
}

func constructArray(n int, k int) []int {
    t := k - 1 // 减去差值为1的

    d := 0
    if t % 2 == 1 {
        d++
    }

    d += t/2

    res := make([]int, n)

    // 一开始的值
    i, end := 0, 0
    if t % 2 == 1 {
        res[i] = n - end
        i++
        end++
    }

    for k, start := true, 1; i < n; i++ {
        // fmt.Println(end, "===>", d)
        if k == true || end >= d {
            k = false
            res[i] = start
            start++
        } else {
            k = true
            res[i] = n-end
            end++
        }
    }

    return res
}
