package main

import (
    "fmt"
    "sort"
)

func main() {
    fmt.Println(minDeletionSize([]string{"ca","bb","ac"}) == 1)
    fmt.Println(minDeletionSize([]string{"xc","yb","za"}) == 0)
    fmt.Println(minDeletionSize([]string{"zyx","wvu","tsr"}) == 3)
    fmt.Println(minDeletionSize([]string{"xga","xfb","yfa"}) == 1)
}


func minDeletionSize(A []string) int {
    l := len(A[0])

    t := 0

    cur := make([]string, len(A))
    cur2:= make([]string, len(A))

    // 每次遍历当前的数据
    for i := 0; i < l; i++ {
        copy(cur, cur2)
        for j := 0; j < len(A); j++ {
            cur[j] = cur[j] + string(A[j][i])
        }

        if k := isSorted(sort.StringSlice(cur)); k {
            copy(cur2, cur)
        } else {
            t++
        }
    }

    return t
}


func isSorted(ss []string) bool {
    for i := 0; i < len(ss) - 1; i++ {
        if ss[i] > ss[i+1] {
            return false
        }
    }

    return true
}