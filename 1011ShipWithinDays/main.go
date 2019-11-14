package main

import (
    "fmt"
)

func main() {
    fmt.Println(shipWithinDays([]int{1,2,3,1,1}, 4))
    fmt.Println(shipWithinDays([]int{1,2,3,4,5,6,7,8,9,10}, 5))
    fmt.Println(shipWithinDays([]int{3,2,2,4,1,4}, 3))
}


func shipWithinDays(weights []int, D int) int {
    l := len(weights)

    // 快速数组
    sum := make([]int, l)
    sum[0] = weights[0]
    mmax := weights[0]
    for i := 1; i < l; i++ {
        sum[i] = sum[i-1] + weights[i]
        mmax = max(mmax, weights[i])
    }

    mmax = max(mmax, sum[l-1]/D)

    for ; ; {

        t, c := 0, 0
        index := -1
        // fmt.Println("======")
        for i := 0; i < l && c < D; {
            index = binarySeach(sum, i, l-1, mmax + t)
            // fmt.Println(index, ",", mmax ,",",mmax + t)
            if index == -1 {
                // 说明没有找到
                break
            }

            t = sum[index]
            i = index + 1
            c++
        }

        // fmt.Println("c", c)
        if c <= D && index != -1 && index >= l - 1{
            return mmax
        }

        mmax++
    }

    return mmax
}

// 查询最后一个不大于target在数组中的位置
func binarySeach(array []int, start, end int, target int) int {
    if array[start] > target {
        return -1
    }

    // 二分查找
    l := start
    for ; start < end; {
        mid := (start+end)/2

        if array[mid] > target {
            end = mid
        } else {
            start = mid + 1
        }
    }

    if array[start] <= target {
        return start
    }

    if end <= l {
        return -1
    }

    return end-1
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
