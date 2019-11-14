package main

import (
    "fmt"
)

func main() {
    fmt.Println(updateMatrix([][]int{{0,0,0},{0,1,0},{0,0,0}}))
    fmt.Println(updateMatrix([][]int{{1,1,1},{1,0,1},{1,1,1}}))
}

func updateMatrix(matrix [][]int) [][]int {
    hight := len(matrix)
    width := len(matrix[0])

    res := make([][]int, hight)
    for i := 0; i < hight; i++ {
        res[i] = make([]int, width)
        for k := range res[i] {
            if matrix[i][k] == 0 {
                res[i][k] = 0
            } else {
                left, up := 100001, 100001
                if k > 0 {
                    left = res[i][k-1]
                }

                if i > 0 {
                    up = res[i-1][k]
                }

                res[i][k] = min(left, up) + 1
            }
        }
    }

    for i := hight-1; i >= 0; i-- {
        for j := width-1; j >= 0; j-- {
            if res[i][j] == 0 {
                continue
            }

            right, down := 100001, 10001
            if j < width -1 {
                right = res[i][j+1]
            }

            if i < hight - 1 {
                down = res[i+1][j]
            }

            res[i][j] = min(res[i][j], min(right, down) + 1)
        }
    }

    return res
}





func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}