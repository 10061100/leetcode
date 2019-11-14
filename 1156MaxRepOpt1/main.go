package main

import (
    "fmt"
)

func main() {
    fmt.Println(maxRepOpt1("ababa") == 3)
    fmt.Println(maxRepOpt1("aaabaaa") == 6)
    fmt.Println(maxRepOpt1("aaabaaabba") == 7)
}


func maxRepOpt1(text string) int {
    count := [26]int{}
    l := len(text)

    // 计数
    for _, v := range text {
        count[v-int32('a')]++
    }

    left := make([]int, l)
    right:= make([]int, l)

    left[0] = 1
    right[l-1] = 1

    res := 0
    c := 1
    for i := 1; i < l; i++ {
        if text[i] == text[i-1] {
            c++
        } else {
            c = 1
        }

        left[i] = c
        res = max(res, c)
    }

    c = 1
    for j := l-2; j >= 0; j-- {
        if text[j] == text[j+1] {
            c++
        } else {
            c = 1
        }

        right[j] = c
        res = max(res, c)
    }

    for i := 1; i < l-1; i++ {
        // 和左边的数据不相同, 且左边的字符在别的地方出现过
        if text[i] != text[i-1] && count[int32(text[i-1]) - int32('a')] > left[i-1] {
            res = max(res, left[i-1] + 1)
        }

        // 如果和右边的数据不相同, 且右边的字符也在别的地方出现过
        if text[i] != text[i+1] && count[int32(text[i+1]) - int32('a')] > right[i+1] {
            res = max(res, right[i+1] + 1)
        }

        // 如果左右的数字相同, 当时和当前的不同
        if text[i-1] == text[i+1] && text[i-1] != text[i] {
            t := count[int32(text[i+1]) - int32('a')]

            // 如果这个字符还在别的地方出现过
            if t > left[i-1] + right[i+1] {
                res = max(res, left[i-1] + right[i+1] + 1)
            } else {
                res = max(res, left[i-1] + right[i+1])
            }
        }
    }

    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
