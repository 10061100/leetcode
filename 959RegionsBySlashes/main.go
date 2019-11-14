package main

import (
    "fmt"
)

func main() {
    fmt.Println(regionsBySlashes([]string{
        " /",
        "/ ",
    }) == 2)
    fmt.Println(regionsBySlashes([]string{
        " /",
        "  ",
    }) == 1)
    fmt.Println(regionsBySlashes([]string{
        "\\/",
        "/\\",
    }) == 4)
    fmt.Println(regionsBySlashes([]string{
        "/\\",
        "\\/",
    }) == 5)
    fmt.Println(regionsBySlashes([]string{
        "//",
        "/ ",
    }) == 3)
    fmt.Println(regionsBySlashes([]string{
        "/",
    }) == 2)
    fmt.Println(regionsBySlashes([]string{
        " ",
    }) == 1)
}


func regionsBySlashes(grid []string) int {

    n := len(grid)
    dp := make([]int, (n + 1) * (n + 1)+1)

    // 初始化并查集
    for i := 0; i <= n; i++ {
        // 上下左右的边都指向0,0 => 0
        dp[index(0, i, n+1)] = 1
        dp[index(i, 0, n+1)] = 1
        dp[index(n, i, n+1)] = 1
        dp[index(i, n, n+1)] = 1
    }
    dp[1] = 0

    res := 1
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == ' ' {
                continue
            }

            a, b := 0, 0
            if grid[i][j] == '\\' {
                a = index(i, j, n+1)
                b = index(i+1, j+1, n+1)
            } else {
                a = index(i+1, j, n+1)
                b = index(i, j+1, n+1)
            }

            aroot := getRoot(a, dp)
            broot := getRoot(b, dp)

            if aroot == broot {
                res++
                continue
            }

            if aroot < broot {
                dp[broot] = aroot
            } else {
                dp[aroot] = broot
            }

            // fmt.Println(dp)
        }

    }

    return res
}


func index(x, y int, n int) int {
    return x * n + y + 1
}


func getRoot(parent int, set []int) int {

    t := parent

    for; set[t] != 0; t = set[t] {}

    // if set[parent] != t {
    //     set[parent] = t
    // }

    return t
}