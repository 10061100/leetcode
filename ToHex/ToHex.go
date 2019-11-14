// ToHex
package main

import (
    "fmt"
)

func main() {

    n := -1//-2147483648
    println(toHex(n))
    fmt.Printf("%b\n", -2 >> 1)
}

func toHex(num int) string {
    if num == 0 {
        return "0"
    }

    neg := false
    if num < 0 {
        neg = true
        num = num & 0x7FFFFFFF
    }

    r := make([]byte, 0)
    for i := 0; i < 7 ; i++ {
        t := num & 0xF
        if neg == false && num != 0 {
            r = append(r, getChar(t))
        }
        if neg == true && num != -1 {
            r = append(r, getChar(t))
        }
        num = num >> 4
    }

    if neg {
        num += 8
    }

    if num != 0 {
        r = append(r, getChar(num))
    }

    return string(convese(r))
}


func getChar(n int) byte {
    switch {
    case n < 0:
        return byte(^n+1) + '0'
    case n < 10:
        return '0' + byte(n)
    default:
        return 'a' + byte(n - 10)
    }
}

func convese(c []byte) []byte {
    l := len(c)
    // fmt.Println(string(c))
    for i := 0; ; i++ {
        op := l - 1 - i
        if op != i {
            // println(i, "=>", op)
            t := c[i]
            c[i] = c[op]
            c[op] = t
        }

        if op - i <= 1 {
            break
        }
    }
    // fmt.Println(string(c))

    return c
}