package main

import (
    "fmt"
)

func main()  {
    fmt.Println(longestCommonSubsequence("abc", "ac") == 2)
    fmt.Println(longestCommonSubsequence("abc", "ebg") == 1)
    fmt.Println(longestCommonSubsequence("abc", "ebc") == 2)
    fmt.Println(longestCommonSubsequence("abc", "efg") == 0)
}

func longestCommonSubsequence(text1 string, text2 string) int {
    l1 := len(text1)
    l2 := len(text2)

    if l1 == 0 || l2 == 0 {
        return 0
    }

    pre := make([]int, l2+1)
    cur := make([]int, l2+1)
    // bt1 := []byte(text1)
    bt2 := []byte(text2)


    for _, v := range []byte(text1) {
        cur[0] = 0
        for i := 1; i <= l2; i++ {
            cur[i] = max(cur[i-1], pre[i])
            if v == bt2[i-1] {
                cur[i] = max(cur[i], pre[i-1]+1)
            }
        }
        copy(pre, cur)
    }

    return pre[l2]
}


func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
