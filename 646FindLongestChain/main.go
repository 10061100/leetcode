package main

import (
    "fmt"
    "sort"
)

func main() {
    fmt.Println(findLongestChain([][]int{{1,2},{2,3},{3,4}}) == 2)
    fmt.Println(findLongestChain([][]int{{1,2},{2,3}}) == 1)
    fmt.Println(findLongestChain([][]int{{1,2},{3,5}, {3,4}, {5,6}}) == 3)
}

func findLongestChain(pairs [][]int) int {
    sort.Slice(pairs, func(i, j int) bool {
        t1 := pairs[i]
        t2 := pairs[j]

        if t1[1] < t2[1] {
            return true
        }

        if t1[1] > t2[1] {
            return false
        }

        if t1[0] <= t2[0] {
            return true
        }

        return false
    })

    t := pairs[0]
    // ti:= 0

    res := 1
    for i := 1; i < len(pairs); i++ {
        if pairs[i][0] <= t[1] {
            continue
        }

        res++
        t = pairs[i]
        //ti = i
    }

    return res
}
