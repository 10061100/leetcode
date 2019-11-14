// RepeatedSubstringPattern
package main

import "fmt"
import "math"

func main() {
    // println(repeatedSubstringPattern2("aba"))
    fmt.Printf("%b\n", 2)
    fmt.Printf("%b\n", ^2 & math.MaxInt32)
    fmt.Printf("%d\n", ^5)
}


func repeatedSubstringPattern(s string) bool {
    bs := []byte(s)
    slen := len(bs)

    if slen == 0 {
        return true
    }

    t := make([]byte, 0)
    t = append(t, bs[0])
    tlen := 1
    cur := 0
    for i, start := 1, 1 ; i < slen ; {
        if t[cur] == bs[i] {
            cur++
            i++
        } else {
            for j := start; j < i ; j++ {
                t = append(t, bs[j])
                tlen++
            }
            if bs[i] != t[0] {
                t = append(t, bs[i])
                tlen++
                i++
            }
            start = i
            cur = 0
        }

        if cur == tlen {
            // 如果t结束了, 就说明肯定是一个轮回了
            cur = 0
        }
    }

    if tlen == slen || cur != 0 {
        return false;
    }

    return true
}


func repeatedSubstringPattern2(s string) bool {
    bs := []byte(s)
    slen := len(bs)

    if slen == 0 {
        return true
    }

    for i := 2; i < slen; i++ {
        if slen % i != 0 {
            continue
        }

        sublen := i
        t := make([]byte, sublen)
        copy(t[0:], bs[0:sublen])
        if checkRepeat(bs, t) {
            return true
        }
    }

    return false
}


func checkRepeat(s , sub []byte ) bool {
    cur := 0
    for i := 0 ; i < len(s); i++{
        if s[i] != sub[cur] {
            return false
        } else {
            cur++
        }

        if cur == len(sub) {
            cur = 0
        }
    }

    return cur == 0
}