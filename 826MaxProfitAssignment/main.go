package main

import (
    "fmt"
    "sort"
)

func main()  {
    fmt.Println(maxProfitAssignment([]int{2,4,6,8,10}, []int{10,20,30,40,50}, []int{4,5,6,7}) == 100)
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
    l := len(difficulty)
    res := 0

    if l <= 0 {
        return res
    }

    jobs := make([]*Job, l)
    for i := 0; i < l; i++ {
        jobs[i] = NewJob(difficulty[i], profit[i])
    }

    sort.Sort(sort.IntSlice(worker))
    sort.Slice(jobs, func(i, j int) bool {
        min := jobs[i]
        max := jobs[j]

        if min.Diff < max.Diff {
            return true
        }

        if min.Profit < min.Profit {
            return true
        }

        return false
    })

    j := 0
    maxProfit := 0
    for _, diff := range worker {
        // 对于当前难度
        for ; j < l; j++ {
            if jobs[j].Diff > diff {
                break
            } else {
                maxProfit = max(maxProfit, jobs[j].Profit)
            }
        }

        res += maxProfit
    }

    return res
}

type Job struct {
    Diff int
    Profit int
}

func NewJob(d , p int) *Job {
    return &Job{
        Diff: d,
        Profit: p,
    }
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}