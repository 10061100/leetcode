package main

import (
    "fmt"
)

func main()  {
    fmt.Println(longestStrChain([]string{"a","b","ba","bca","bda","bdca"}) == 4)
    fmt.Println(longestStrChain([]string{"a","b"}) == 1)
    fmt.Println(longestStrChain([]string{"a","b","bca","bda","bdca"}))
}

func longestStrChain(words []string) int {
    edges := buildMap(words)

    res := make([]int, len(words))

    m := 0
    for i := 0; i < len(res); i++ {
        if res[i] == 0 {
            dfs(i, edges, words, res)
        }

        m = max(m, res[i])
    }



    return m
}

func dfs(i int, edges[][]int, word []string, res []int) int {
    if res[i] == 0 {
        t := 0
        for _, v := range edges[i] {
            t = max(t, dfs(v, edges, word, res))
        }

        res[i] = 1 + t
    }

    return res[i]
}

func max(a, b int) int {
    if a > b{
        return a
    }

    return b
}

func buildMap(words []string) [][]int {
    res := make([][]int, len(words))
    for i := 0; i < len(words); i++ {
        res[i] = make([]int, 0)
    }
    for i := 0; i < len(words); i++ {
        for j := i+1; j < len(words); j++ {
            parent, _, issub := isSubSeq(words[i], words[j])
            if !issub {
                continue
            }

            parentIdex , sonIdex:= i, j
            if parent != words[i] {
                parentIdex = j
                sonIdex = i
            }


            res[parentIdex] = append(res[parentIdex], sonIdex)
        }
    }

    return res
}

func isSubSeq(parent, son string) (nparent, nson string, issub bool) {
    if len(son) == len(parent) {
        return "", "", false
    }

    if len(son) > len(parent) {
        parent, son = son , parent
    }

    if len(son) + 1 != len(parent) {
        return "", "", false
    }

    i, j := 0, 0

    for ; i < len(son) && j < len(parent); {
        if son[i] == parent[j] {
            i++
            j++
            continue
        }

        j++
    }

    if i == len(son) && j == len(parent) {
        return parent, son, true
    }

    if i == len(son) && j == len(parent)-1 {
        return parent, son, true
    }

    return "", "", false
}