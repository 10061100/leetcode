package main

import (
    "fmt"
    "sort"
)

func main() {
    fmt.Println(deckRevealedIncreasing([]int{17,13,11,2,3,5,7}))
}

func deckRevealedIncreasing(deck []int) []int {
    sort.Sort(sort.IntSlice(deck))
    res := make([]int, len(deck))

    cur := 0
    l   := len(deck)
    for _, v := range deck {
        res[cur%l] = v
        cur += 2
    }

    return res
}
