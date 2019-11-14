package main

import (
    "fmt"
)

func main() {
    fmt.Println(customSortString("cba", "abcd"))
}


func customSortString(S string, T string) string {
    m := [26]int{}

    for _, v := range T {
        m[int(v) - int('a')]++
    }

    res := make([]byte, len(T))
    i   := 0
    for _, v := range S {
        t := int(v) - int('a')
        if m[t] > 0 {
            for ; m[t] > 0; m[t]-- {
                res[i] = byte(v)
                i++
            }
        }
    }

    for k, v := range m {
        if v > 0 {
            for ; m[k] > 0 ; m[k]-- {
                res[i] = byte(k + 'a')
                i++
            }
        }
    }

    return string(res)
}

