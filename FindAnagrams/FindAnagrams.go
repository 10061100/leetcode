// FindAnagrams
package main

import (
    "fmt"
)

func main() {
    fmt.Println(findAnagrams("abab", "ab"))
}

func findAnagrams(s string, p string) []int {
    r := make([]int, 0)
    pmap := make(map[byte]int)
    for _, v := range []byte(p) {
        if _, ok := pmap[v]; !ok {
            pmap[v] = 1
        } else {
            pmap[v] += 1
        }
    }

    lp := len(p)
    ls := len(s)
    bs := []byte(s)
    nozeroCount := len(pmap)

    for i, j := 0, 0; i < ls; i++ {
        c := bs[i]
        if _, ok := pmap[c]; ok {
            pmap[c] -= 1
            if pmap[c] == 0 {
                nozeroCount--
            } else if pmap[c] == -1 {
                nozeroCount++
            }
        }

        if i >= lp {
            jc := bs[j]
            if _, ok := pmap[jc]; ok {
                pmap[jc] += 1
                if pmap[jc] == 0 {
                    nozeroCount--
                } else if pmap[jc] == 1 {
                    nozeroCount++
                }
            }
            j++
        }

        if nozeroCount == 0 {
            r = append(r, j)
        }
    }

    return r
}