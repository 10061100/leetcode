package main

import (
    "fmt"
)

const F = 1000000007

func main()  {
    var t, k int

    fmt.Scanf("%d %d", &t, &k)

    input := make([][]int, t)
    l := 0
    for i := 0; i < t; i++ {
        input[i] = make([]int, 2)
        fmt.Scanf("%d %d", &input[i][0], &input[i][1])
        l = max(l, input[i][1])
    }

    dp1 := make([]int, l+1) // 红花
    dp2 := make([]int, l+1) // 百花
    sum := make([]int, l+1) // 和

    dp1[0] = 0
    dp1[1] = 1
    for i := 0; i < k; i++ {
        dp2[i] = 0
    }
    dp2[k] = 1

    sum[0] = 0
    sum[1] = dp1[1] + dp2[1] // 和

    for i := 2; i <= l; i++ {
        dp1[i] = (dp2[i-1] + dp1[i-1]) % F
        if i > k {
            dp2[i] = (dp2[i-k] + dp1[i-k])%F
        }

        sum[i] = (sum[i-1] + (dp1[i] + dp2[i]) % F) % F
    }


    for i := 0; i < t ; i++ {
        m := sum[input[i][1]]
        if input[i][0] > 0 {
            m = m - sum[input[i][0]-1]
        }

        fmt.Printf("%d\n", (m+F)%F)
    }
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
