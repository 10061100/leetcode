package main

import (
    "fmt"
)

func main() {
    fmt.Println(numTilePossibilities("AB"))
}


func numTilePossibilities(tiles string) int {
    t := make([]int, 26)
    for _, v := range []byte(tiles) {
        t[v-'A']++
    }

    cache := make(map[string]int)
    cache[""] = 1
    res := 0

    dfs(t, cache, &res)

    return res
}


func dfs( t []int , cache map[string]int, res *int) int {
    key := buildkey(t)
    if _, ok := cache[key]; !ok {
        m := 0
        for i := 0; i < 26; i++ {
            if t[i] > 0 {
                t[i]--
                m += dfs(t, cache, res)
                t[i]++
            }
        }
        cache[key] = m
        (*res) += m
    }

    return cache[key]
}

func buildkey(t []int) string {
    m := make([]byte, 0)
    for i := 0; i < 26; i++ {
        if t[i] > 0 {
            m = append(m, byte(i) + 'A', byte(t[i]))
        }
    }

    return string(m)
}