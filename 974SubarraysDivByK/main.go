package main

import (
    "fmt"
)

func main() {
    fmt.Println(subarraysDivByK([]int{4,5,0,-2,-3,1}, 5) == 7)
}

func subarraysDivByK(A []int, K int) int {
    sum := make([]int, len(A))
    sum[0] = A[0]
    tmp := make([]int, K)
    tmp[((sum[0]%K) + K)%K]++

    for i := 1; i < len(A); i++ {
        sum[i] = sum[i-1] + A[i]
        tmp[(sum[i]%K + K) % K]++
    }

    res := tmp[0]
    for _, v := range tmp {
        res += v * (v-1)/2
    }

    return res
}

