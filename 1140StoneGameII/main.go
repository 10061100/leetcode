package main

import (
    "fmt"
    "strconv"
)

func main() {
    fmt.Println(stoneGameII([]int{2,7,9,4,4}) == 10)
    fmt.Println(stoneGameII([]int{7,9,4,4}) == 16)
    fmt.Println(stoneGameII([]int{4,4,7,9}) == 13)
}

func stoneGameII(piles []int) int {
    sum := make([]int, len(piles) + 1)
    sum[len(piles)] = 0
    for i := len(piles) - 1; i >= 0; i-- {
        sum[i] = sum[i+1] + piles[i]
    }

    cache := make(map[string]int)

    res := dfs(0, 1, cache, piles, sum)

    return res
}


func dfs(start, m int, cache map[string]int, piles []int, sum []int) int {

    if start >= len(piles) {
        return 0
    }

    k := strconv.Itoa(start) + "-" + strconv.Itoa(m)

    // 如果已经计算了, 就直接返回计算的值
    if v, ok := cache[k]; ok {
        return v
    }

    // 进行计算, cache里面存的是, 从start堆开始, 一次最多拣选m的数量时 能拣选的最多值
    res := 0
    mmax:= 0
    for i := 0; i < 2*m && start + i < len(piles); i++ {
        res += piles[start+i]
        next := dfs(start+i+1, max(m, i+1), cache, piles, sum) // 这个是下一个人最多能获取的数量
        t    := sum[start+i+1] - next // 这个事当前的人后面最多拣选的数量
        if tmp := res + t; tmp > mmax {
            mmax = tmp
        }
    }

    cache[k] = mmax

    return mmax
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
