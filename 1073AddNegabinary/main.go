package main

import (
    "fmt"
)

func main() {
    fmt.Println(addNegabinary([]int{1,1,1,1,1}, []int{1,0,1}))
    fmt.Println(addNegabinary([]int{1,0,1}, []int{1,0,1}))
}

func addNegabinary(arr1 []int, arr2 []int) []int {
    res := make([]int, 0)

    for i, j := len(arr1)-1, len(arr2)-1; i >= 0 || j >= 0; {

        if i >= 0 && j >= 0 {
            res = append(res, arr1[i] + arr2[j])
            i--
            j--
        } else if i >= 0 {
            res = append(res, arr1[i])
            i--
        } else {
            res = append(res, arr2[j])
            j--
        }
    }

    // 从头开始计算, 每当n位需要进1的时候, n+1 和 n+2 同时加1
    l := len(res)
    for i := 0 ; i < l ; i++ {
        if res[i] == 1 || res[i] == 0 {
            continue
        }



        for ; res[i] > 1 ; {
            if i == l-1 {
                res = append(res, 0, 0)
                l+=2
            } else if i == l - 2 {
                res = append(res, 0)
                l++
            }

            if res[i+1] > 0 {
                res[i] -= 2
                res[i+1] -= 1
            } else {
                res[i] -= 2
                res[i+1] += 1
                res[i+2] += 1
            }
        }
    }

    // fmt.Println(res)
    newl := l-1
    // fmt.Println(newl, ",", l)
    for ; newl >= 0 && res[newl] == 0; newl-- {} // 过滤掉前驱的0

    for i, j := 0, newl; i < j; i, j = i+1, j-1 {
            res[i], res[j] = res[j], res[i]
    }

    if newl < 0 {
        newl = 0
    }

    return res[0:newl+1]
}
