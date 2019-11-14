package main

import (
    "fmt"
)

func main() {
    fmt.Println(isIdealPermutation([]int{0,2,1}) == true)
    fmt.Println(isIdealPermutation([]int{1,0,2}) == true)
    fmt.Println(isIdealPermutation([]int{1,2,0}) == false)
    fmt.Println(isIdealPermutation([]int{1,0}) == true)
    fmt.Println(isIdealPermutation([]int{0,1}) == true)
    fmt.Println(isIdealPermutation([]int{1}) == true)
}


func isIdealPermutation(A []int) bool {
    l := len(A)

    if l <= 2 {
        return true
    }

    min1, min2, min3 := A[1], A[0], -1
    preIsBig := false
    if min2 > min1 {
        preIsBig = true
    }

    for i := 2; i < l; i++ {
        if A[i] > A[i-1] {
            if A[i] <= min1 || A[i] <= min2 || A[i] <= min3 {
                return false
            }
            preIsBig = false
        } else {
            if preIsBig {
                return false
            }
            if A[i] <= min2 || A[i] <= min3 {
                return false
            }
            preIsBig = true
        }

        min1, min2, min3 = A[i], min1, min2
    }

    return true
}

