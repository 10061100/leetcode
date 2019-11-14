package main

import (
    "fmt"
    "sort"
)

func main() {
    fmt.Println(smallestRangeII([]int{1,3,6}, 3) == 3)
    fmt.Println(smallestRangeII([]int{0,10}, 2) == 6)
    fmt.Println(smallestRangeII([]int{1}, 0) == 0)
    fmt.Println(smallestRangeII([]int{7,8,8}, 5) == 1)
    fmt.Println(smallestRangeII([]int{3,4,7,0}, 5) == 7)
    fmt.Println(smallestRangeII([]int{7,8,8,5,2}, 4) == 5)
}


// 这里就是找一个点, 所有点左边的(包括点)进行+K, 所有右边的进行-K, 这样值最小
func smallestRangeII(A []int, K int) int {
    sort.Sort(sort.IntSlice(A))
    l := len(A)

    res := A[l-1] - A[0]
    for i := 0; i < l-1; i++ {
        tmin := min(A[0]+K, A[i+1]-K)
        tmax := max(A[i]+K, A[l-1]-K)
        res  = min(res, tmax - tmin)
    }

    return res
}

func min( a, b int) int {
    if a > b {
        return b
    }

    return a
}


func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}

