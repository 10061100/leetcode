package main

import (
    "fmt"
)

func main() {
    fmt.Println(minSteps(6))
}


func minSteps(n int) int {
    cache := make(map[int]int)
    cache[1] = 0
    return dfs(n, cache)
}


func dfs(n int, cache map[int]int) int {
    if v, ok := cache[n]; ok {
        return v
    }

    min := n

    for i := 2; i * i <= n; i++ {
        if n % i != 0 {
            continue
        }

        t1 := i
        t2 := dfs(n/i, cache)

        if t := t1 + t2; t < min {
            min = t
        }
    }

    cache[n] = min

    return min
}
