package main

import (
    "fmt"
    "sort"
)

func main() {
    // fmt.Println(intersectionSizeTwo([][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}}))
    // fmt.Println(intersectionSizeTwo([][]int{{1, 3}, {1, 4}, {2, 5}, {3, 5}}))
    fmt.Println(intersectionSizeTwo([][]int{{2,10}, {3,7}, {3,15}, {4,11}, {6,12}, {6,16}, {7,8}, {7,11}, {7,15}, {11,12}}))
    fmt.Println(intersectionSizeTwo([][]int{{12,19},{18,25},{4,6},{19,24},{19,22}}))
}

func intersectionSizeTwo(intervals [][]int) int {
    sort.Sort(Slice(intervals))

    end := intervals[0][1]
    start := intervals[0][1] - 1
    res := 2
    // fmt.Println(start, ",", end)

    for i := 1; i < len(intervals); i++ {
        newend   := intervals[i][1]
        newstart := intervals[i][0]

        if newstart > end {
            end = newend
            start = end - 1
            res += 2
        } else if newstart == end || newstart > start {
            start = end
            end = newend
            res += 1
        }
    }

    // fmt.Println(end, start)
    return res
}


type Slice [][]int

func( s Slice) Len() int {
    return len(s)
}

func (s Slice) Less(i, j int) bool {
    if s[i][1] < s[j][1] {
        return true
    }

    if s[i][1] == s[j][1] {
        return s[i][0] > s[j][0]
    }

    return false
}

func (s Slice) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}


func min( a, b int) int {
    if a < b {
        return a
    }

    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}