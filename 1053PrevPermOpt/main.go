package main

import (
    "fmt"
)

func main()  {
    fmt.Println(prevPermOpt1([]int{3,2,1}))
    fmt.Println(prevPermOpt1([]int{1,1,5}))
    fmt.Println(prevPermOpt1([]int{1,9,4,6,7}))
    fmt.Println(prevPermOpt1([]int{3,1,1,3}))
}


func prevPermOpt1(A []int) []int {

    n := len(A)

    // 从尾向头遍历,找到第一个序列对, A[i] > A[i+1], 说明i和i+1可以进行交换, 形成一个比较大的符合要求的序列
    // 然后从i+1开始往尾遍历, 找到比A[i+1]
    for i := n-2; i >= 0; i-- {
        if A[i] <= A[i+1] {
            continue
        }

        t := A[i]
        // A[i] = A[i+1]

        m := A[i+1]
        mi:= i+1
        for j := i+2; j < n; j++ {
            if A[i] > A[j] && A[j] > m {
                m = A[j]
                mi = j
            }
        }

        A[i] = m
        A[mi] = t

        break
    }


    return A
}
