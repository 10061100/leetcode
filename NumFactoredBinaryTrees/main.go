package main

import (
    "sort"
)

const MOD = 1000000000 + 7

func numFactoredBinaryTrees(A []int) int {
    l := len(A)

    if l == 1 {
        return 1
    }

    if l == 0 {
        return 0
    }

    sort.Sort(sort.IntSlice(A))
    m := make(map[int]int)
    m[A[0]] = 1
    r := 1
    for i := 1; i < l ; i++ {
        tr := 1
        for j := 0; j < i ; j++ {
            // 能够整除法
            if A[i] % A[j] == 0 {
                t := A[i]/A[j]
                if v, ok := m[t]; ok {
                    n := int((int64(m[A[j]]) * int64(v)))%MOD
                    tr = (tr + n) % MOD
                }
            }
        }
        m[A[i]] = tr
        r = (r + tr) % MOD
    }


    return r
}
