package main

import (
    "fmt"
)

func main() {
    fmt.Println(longestOnes([]int{0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1}, 3))
    fmt.Println(longestOnes([]int{1,1,1,0,0,0,1,1,1,1,0}, 2))
    fmt.Println(longestOnes([]int{1,1,1,0,0,0,1,1,1,1,0}, 3))
}



func longestOnes(A []int, K int) int {
    start := 0
    end   := 0
    t     := 0 //

    res := 0
    for ;  end < len(A); end++ {
        // total := end - start + 1 // 总共有多少数字
        if A[end] == 0 {
            t++ // 如果是一个0的话, 需要占用名额
        }

        for ; t > K;  start++ {
            if A[start] == 0 {
                t--
            }
        }

        res = max(res, end - start + 1)
    }

    return res
}


func max(a , b int) int {
    if a > b {
        return a
    }

    return b
}
