package main

import (
    "fmt"
)

func main() {
    // [1,4,2,2,3,4,3,1]
    // fmt.Println(removeBoxes([]int{1, 3, 2, 2, 2, 3, 4, 3, 1}))
    fmt.Println(removeBoxes([]int{2, 2,1}))
}


func removeBoxes(boxes []int) int {
    l := len(boxes)
    dp := make([][][]int, l)

    // 初始化dp, 数字左边填充相同的数字
    for i := 0; i < l; i++ {
        dp[i] = make([][]int, l)
        for j := 0; j < l; j++ {
            dp[i][j] = make([]int, l)
        }

        for k := 0; k <= i; k++ {
            dp[i][i][k] = (k+1)*(k+1)
        }
    }

    for length := 2; length <= l; length++ {
        for i := 0; i + length <=l; i++ {
            j := i + length - 1
            for k := 0; k <= i; k++ {
                dp[i][j][k] = (k+1) * (k+1) + dp[i+1][j][0]
                for cur := i+1; cur <= j; cur++ {
                    if boxes[cur] != boxes[i] {
                        continue
                    }

                    pre := 0
                    if cur > i+1 {
                        pre = dp[i+1][cur-1][0]
                    }

                    pre = pre + dp[cur][j][k+1]

                    if dp[i][j][k] == 0 || dp[i][j][k] < pre {
                        dp[i][j][k] = pre
                        // fmt.Println(pre)
                    }
                }
            }
        }
    }

    return dp[0][l-1][0]
}
