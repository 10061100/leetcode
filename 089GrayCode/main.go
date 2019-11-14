package main

import (
    "fmt"
)

func main() {
    fmt.Println(grayCode(1))
    fmt.Println(grayCode(2))
}


func grayCode(n int) []int {
    res := make([]int, 0)

    if n == 0 {
        return res
    }

    res = append(res, 0)
    res = append(res, 1)
    t := 1

    for i := 1; i < n; i++ {
        t = t << 1
        l := len(res) - 1
        for ; l >= 0; l-- {
            res = append(res, t + res[l])
        }
    }

    return res
}

