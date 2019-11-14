package main

import (
    "fmt"
    "sort"
)

func main()  {
    var n int
    fmt.Scanf("%d", &n)

    time := make([][]int, n)

    for i := 0; i < n; i++ {
        time[i] = make([]int, 2)
        fmt.Scanf("%d %d", &time[i][0], &time[i][1])
    }

    sort.Slice(time, func(i, j int) bool {
        if time[i][0] < time[j][0] {
            return true
        }

        if time[i][1] < time[j][1] {
            return true
        }

        return false
    })

    r1 := -1
    r2 := -1
    for i := 0; i < n; i++ {
        if time[i][0] > r1 {
            r1 = time[i][1]
            continue
        }

        if time[i][0] > r2 {
            r2 = time[i][1]
            continue
        }

        fmt.Println("NO")
        return
    }

    fmt.Println("YES")
}
