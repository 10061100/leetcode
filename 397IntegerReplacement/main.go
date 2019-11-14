package main

import (
    "fmt"
)

func main()  {
    // fmt.Println(integerReplacement(7))
    // fmt.Println(integerReplacement(8))
    fmt.Println(integerReplacement(6))
}


func integerReplacement(n int) int {
    cache := make(map[int]int)
    cache[1] = 0
    return dfs(n, cache)
}

func dfs(n int, cache map[int]int) int {
    if _, ok := cache[n]; !ok {
        if n % 2 == 0 {
            cache[n] = 1 + dfs(n >> 1, cache)
        } else {
            cache[n] = 1 + min(dfs(n-1, cache), dfs(n+1, cache))
        }
    }

    return cache[n]
}

func min(a, b int) int {
    if a > b {
        return b
    }

    return a
}


