package main

import (
    "fmt"
)

func main()  {
    fmt.Println(loudAndRich([][]int{{1,0},{2,0}}, []int{1,0,2}))
}


func loudAndRich(richer [][]int, quiet []int) []int {
    edges := make([][]int, len(quiet))

    // 遍历边
    for i := 0; i < len(richer); i++ {
        richone := richer[i][0]
        poorone := richer[i][1]

        if edges[poorone] == nil {
            edges[poorone] = make([]int, 0)
        }

        edges[poorone] = append(edges[poorone], richone)
    }

    res := make([]int, len(quiet))
    for i := 0; i < len(res); i++ {
        res[i] = -1
    }

    for k, v := range res {
        if v == -1 {
            dfs(k, quiet, res, edges)
        }
    }

    return res
}


func dfs(node int, quiet []int, res []int, edges [][]int) int {

    if res[node] == -1 {
        t := node
        for _, v := range edges[node] {
            t = min(quiet[t], dfs(v, quiet, res, edges))
        }

        res[node] = t
    }


    return res[node]
}


func min(a, b int) int {
    if a > b {
        return b
    }

    return a
}

