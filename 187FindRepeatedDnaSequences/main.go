package main

import (
    "fmt"
)

func main() {
    fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
    fmt.Println(findRepeatedDnaSequences(""))
}

func findRepeatedDnaSequences(s string) []string {
    N  := 10
    count := make(map[int]int)
    table := make(map[int]string)
    for i := 0; i + N <= len(s); i++ {
        t := s[i:i+N]
        hash := getHash(t)
        if _, ok := count[hash]; !ok {
            count[hash] = 0
        }
        count[hash]++

        if _, ok := table[hash]; !ok {
            table[hash] = t
        }
    }

    res := make([]string, 0)
    for hash, c := range count {
        if c > 1 {
            res = append(res, table[hash])
        }
    }

    return res
}

func getHash(str string) int {
    k := 0
    for _, v := range []byte(str) {
        n := getIntVal(byte(v)) - 1
        k = k * 4 + n
    }

    return k
}

func getIntVal(s byte) int {
    switch s {
    case 'A': return 1
    case 'C': return 2
    case 'G': return 3
    case 'T': return 4
    }

    return 4
}

