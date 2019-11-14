package main

import (
    "fmt"
    "sort"
)

func main() {
    fmt.Println(rearrangeBarcodes([]int{2,2,1,3}))
}

type Count struct {
    Val int
    C   int
}

func NewCount(v int, c int) *Count {
    return &Count{
        Val: v,
        C: c,
    }
}

func rearrangeBarcodes(barcodes []int) []int {
    m := [10001]int{}

    for _, v := range barcodes {
        m[v]++
    }

    counts := make([]*Count, 0)
    for i := 0; i < 10001; i++ {
        if m[i] == 0{
            continue
        }

        counts = append(counts, NewCount(i, m[i]))
    }


    // 排序
    sort.Slice(counts, func(i, j int) bool {
        if counts[i].C >= counts[j].C {
            return true
        }
        return false
    })

    res := make([]int, len(barcodes))
    i := 0
    for j := 0; j < len(counts); j++ {
        for ; counts[j].C > 0; counts[j].C-- {
            res[i] = counts[j].Val
            i += 2
            if i >= len(barcodes) {
                i = 1
            }
        }
    }

    return res
}

