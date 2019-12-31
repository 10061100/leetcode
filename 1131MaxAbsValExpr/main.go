package main

import (
    "fmt"
)

func main() {
    fmt.Println(maxAbsValExpr([]int{1,2,3,4}, []int{-1,4,5,6}) == 13)
    fmt.Println(maxAbsValExpr([]int{1,-2,-5,0,10}, []int{0,-2,-1,-7,-4}) == 20)
}

func maxAbsValExpr(arr1 []int, arr2 []int) int {
    mmax := []int{-100000000, -100000000, -100000000, -100000000}
    mmin := []int{100000000,100000000,100000000,100000000}

    for i := 0; i < len(arr1); i++ {
        mmax[0] = max(mmax[0], arr1[i] + arr2[i] - i)
        mmax[1] = max(mmax[1], arr1[i] - arr2[i] - i)
        mmax[2] = max(mmax[2], arr2[i] - arr1[i] - i)
        mmax[3] = max(mmax[3], arr1[i] + arr2[i] + i)


        mmin[0] = min(mmin[0], arr1[i] + arr2[i] - i)
        mmin[1] = min(mmin[1], arr1[i] - arr2[i] - i)
        mmin[2] = min(mmin[2], arr2[i] - arr1[i] - i)
        mmin[3] = min(mmin[3], arr1[i] + arr2[i] + i)

        // fmt.Println("========")
        // fmt.Println(mmax)
        // fmt.Println(mmin)
    }

    ret := -1
    for i := 0; i < 4 ; i++ {
        ret = max(ret, mmax[i] - mmin[i])
    }

    return ret
}


func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}


func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}