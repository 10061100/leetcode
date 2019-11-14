package main

import (
    "fmt"
)

func main() {
    fmt.Println(isMatch("aa", "a."))
    fmt.Println(isMatch("aa", "a*"))
    fmt.Println(isMatch("ab", "c*a*b*"))
    fmt.Println(isMatch("ab", ".*"))
}


type Char struct {
    C byte
    Count int
}

func NewChar(c byte, n int) *Char {
    return &Char{
        C :c,
        Count: n,
    }
}

func isMatch(s string, p string) bool {
    if s == p {
        return true
    }

    bs := []byte(s)
    bp := []byte(p)

    ls := len(bs)
    lp := len(bp)

    char := make([]*Char, 0)
    for i := 0; i < lp; i++{
        if i == lp - 1 {
            char = append(char, NewChar(bp[i],1))
            continue
        }

        if bp[i+1] == '*' {
            char = append(char, NewChar(bp[i], -1))
            i++
            continue
        }

        char = append(char, NewChar(bp[i], 1))
    }

    lp = len(char)
    dp := make([][]bool, ls+1)
    for i := 0; i <= ls; i++ {
        dp[i] = make([]bool, lp+1)
    }

    dp[0][0] = true
    for i := 1; i <= ls; i++ {
        dp[i][0] = false
    }

    for i := 1; i <= lp; i++ {
        c := char[i-1]
        if c.Count == -1 {
            dp[0][i] = dp[0][i-1]
        } else {
            dp[0][i] = false
        }
    }

    for i := 1; i <= ls; i++ {
        for j := 1; j <= lp; j++ {
            curs := bs[i-1]
            curp := char[j-1]

            if curp.Count == -1 {
                if curp.C == '.' {
                    dp[i][j] = dp[i-1][j-1] || dp[i][j-1] || dp[i-1][j]
                } else {
                    dp[i][j] = dp[i][j-1] || ((curp.C == curs) && (dp[i-1][j-1] || dp[i-1][j]) )
                }

            } else {
                // 如果是一个数字
                if curp.C == '.'  || curs == curp.C{
                    dp[i][j] = dp[i-1][j-1]
                } else {
                    dp[i][j] = false
                }
            }
        }
    }

    // fmt.Println(dp)
    return dp[ls][lp]
}


