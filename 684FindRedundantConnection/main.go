package main

import (
    "fmt"
)

func main() {
    fmt.Println(findRedundantConnection([][]int{{1,2},{1,3},{2,3}}))
    fmt.Println(findRedundantConnection([][]int{{1,2}, {2,3}, {3,4}, {1,4}, {1,5}}))
    fmt.Println(findRedundantConnection([][]int{{1,4},{3,4},{1,3},{1,2},{4,5}}))
    fmt.Println(findRedundantConnection([][]int{{3,4},{1,2},{2,4},{3,5},{2,5}}))
}

func findRedundantConnection(edges [][]int) []int {
    l := len(edges)
    set := make([]int, l+1)

    res := []int{0,0}
    for _, v := range edges {
        parent := v[0]
        son    := v[1]

        parentRoot := getRoot(set, parent)
        sonRoot    := getRoot(set, son)

        // 是一个root, 则说明这个边是重复的
        if sonRoot == parentRoot {
            res = []int{parent, son}
            break
        }

        if parentRoot < sonRoot {
            set[sonRoot] = parentRoot
        } else {
            set[parentRoot] = sonRoot
        }
    }

    return res
}

func getRoot(set []int, parent int) int {

    for ; set[parent] != 0 ; {
        parent = set[parent]
    }

    return parent
}
