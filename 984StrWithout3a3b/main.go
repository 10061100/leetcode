package main

import (
    "fmt"
)

func main() {
    fmt.Println(strWithout3a3b(2,0))
    fmt.Println(strWithout3a3b(0,2))
    fmt.Println(strWithout3a3b(3,1))
    fmt.Println(strWithout3a3b(4,1))
    fmt.Println(strWithout3a3b(4,3))
}


func strWithout3a3b(A int, B int) string {
    res := make([]byte, 0)
    a := byte('a')
    b := byte('b')

    if A < B {
        a, b = b, a
        A, B = B, A
    }

    for ; A != 0 && B != 0; {
        // 如果a比较多, 就两个a, 一个b
        if A > B {
            res = append(res, a, a, b)
            A -= 2
            B -= 1
        } else {
            res = append(res, a, b)
            A--
            B--
        }
    }

    t := A
    if B != 0 {
        t = B
        a = b
    }

    for ; t > 0; t-- {
        res = append(res, a)
    }


    return string(res)
}
