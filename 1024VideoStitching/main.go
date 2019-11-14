package main

import (
    "fmt"
    "sort"
)

func main() {
    fmt.Println(videoStitching([][]int{{0,1}, {1,2}}, 7) == -1 )
    fmt.Println(videoStitching([][]int{{0,1},{6,8},{0,2},{5,6},{0,4},{0,3},{6,7},{1,3},{4,7},{1,4},{2,5},{2,6},{3,4},{4,5},{5,7},{6,9}}, 9) == 3)
    fmt.Println(videoStitching([][]int{{0,4}, {2,8}}, 8) == 2)
    fmt.Println(videoStitching([][]int{{0,0},{9,9},{2,10},{0,3},{0,5},{3,4},{6,10},{1,2},{4,7},{5,6}}, 5))
}

func videoStitching(clips [][]int, T int) int {
    sort.Sort(Slice(clips))
    start := 0
    count := 0
    // fmt.Println(clips)
    for i := 0; i < len(clips) && start < T ; {
        mmax := -1
        for ; i < len(clips) && clips[i][0] <= start; i++ {
            mmax = max(mmax, clips[i][1])
        }

        if mmax == -1 {
            // 说明找不到
            return -1
        }
        count++
        start = mmax
    }

    if start < T {
        return -1
    }

    return count
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}


type Slice [][]int

func (s Slice) Len() int {
    return len(s)
}

func (s Slice) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s Slice) Less(i, j int) bool {
    if s[i][0] < s[j][0] {
        return true
    }

    if s[i][0] > s[j][0] {
        return false
    }

    if s[i][1] >= s[j][1] {
        return true
    }

    return false
}