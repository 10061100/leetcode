package main

import (
    "fmt"
)

func main()  {
    fmt.Println(pathInZigZagTree(14))
}

func pathInZigZagTree(label int) []int {
    res := make([]int, 0)

    for ; label != 0; {
        res = append(res, label)
        label /= 2
    }

    for i, j := 0, len(res)-1; i <= j; i, j = i+1, j-1 {
        res[i], res[j] = res[j], res[i]
    }

    //fmt.Println(res)
    if len(res) <= 1 {
        return res
    }

    t := 2
    pre := 1
    m := len(res)%2
    for k := 1; k < len(res); k++ {
        cur := t * (t * 2 - 1)
        sum := cur - pre
        // fmt.Println(cur, " => ", sum, " => ", t, " => ", k)
        if k % 2 == m {
            res[k] = sum/(t/2) - res[k]
        }
        pre = cur
        t = t * 2
    }

    return res
}