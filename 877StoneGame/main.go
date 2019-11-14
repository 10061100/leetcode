package main

import (
    "fmt"
)

func main()  {

    fmt.Println(stoneGame([]int{5,3,4,5}))
}


func stoneGame(piles []int) bool {

    l := len(piles)

    dp1 := make([][]int, l) // 第一个人
    dp2 := make([][]int, l) // 第二个人

    for i := 0; i < l ; i++ {
        dp1[i] = make([]int, l+1)
        dp2[i] = make([]int, l+1) // 右边是开区间
    }

    for i := 0; i < l; i++ {
        dp1[i][i+1] = piles[i] // 如果只有一堆石子, 只能第一个人拿
    }

    for i := 2; i <= l; i++ {
        for j := 0; j + i <= l; j++ {
            // 第一个人他会有两个选择, 最左边和最右边, 选取其中一个他拿的多的
            // [j, j+i-1]
            t1 := piles[j] + dp2[j+1][j+i]
            t2 := piles[j+i-1] + dp2[j][j+i-1]

            if t1 > t2 {
                // 说明第一手拿的值比较多
                dp1[j][j+i] = t1
                dp2[j][j+i] = dp1[j+1][j+i]
            } else {
                dp1[j][j+i] = t2
                dp2[j][j+i] = dp1[j][j+i-1]
            }
        }
    }

    return dp1[0][l] > dp2[0][l]
}
