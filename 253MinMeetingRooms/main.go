package main

import (
    "sort"
)

func main() {
    println(minMeetingRooms([][]int{{0, 30},{5, 10},{15, 20}}))
    println(minMeetingRooms([][]int{{7,10},{2,4}}))
}

func minMeetingRooms(intervals [][]int) int {
    sort.Sort(Intervals(intervals))

    t := make([]int, 0)
    for _, v := range intervals {
        start := v[0]
        end   := v[1]

        hascount := false
        for k, e := range t {
            if e <= start {
                t[k] = end
                hascount = true
                break
            }
        }

        if hascount == false {
            t = append(t, end)
        }
    }

    return len(t)
}



type Intervals [][]int

func (s Intervals) Len() int {
    return len(s)
}

func (s Intervals) Swap(i, j int){
    s[i], s[j] = s[j], s[i]
}

func (s Intervals) Less(i, j int) bool {
    if s[i][0] < s[j][0] {
        return true
    }

    if s[i][0] > s[j][0] {
        return false
    }

    return s[i][1] < s[j][1]
}

