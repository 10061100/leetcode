package main

import (
    "sort"
)

func numMovesStonesII(stones []int) []int {
    n := len(stones)

    sort.Sort(sort.IntSlice(stones))

    // 从最后的结果去看, 一定是连续的一个石头序列
    // 最大的做法一定<= 把中间空闲的点都碰一遍
    // 第一步无论如何都需要挪动0或者n-1, 如果挪动0的话, 呢0和1之间的就没法都碰一次了, 同理n-1也一样
    // 后面就可以保证每次两边端点都+1和-1
    mx := stones[n-1] - stones[0] + 1 - n
    mx = mx - min(stones[1]-stones[0], stones[n-1] - stones[n-2]) - 1

    // 最小值的话用滑动窗口来看
    // 一定至少有一个点是没有移动过的
    mn := mx
    for i, j := 0, 0; i < n; i++ {
        // 每次往后找找到当前比 stones[i] + n - 1 的数
        for ; j + 1 < n && stones[j+1] <= stones[i] + n - 1; j++ {}

        cost := n - (j+1-i)
        // 判断是不是边界 1234,7
        if j - i + 1 == n - 1 && stones[j] - stones[i] + 1 == n-1 {
            cost = 2
        }

        mn = min(mn, cost)
    }

    return []int{mn,mx}
}


func min(a , b int) int {
    if a >b {
        return b
    }

    return a
}