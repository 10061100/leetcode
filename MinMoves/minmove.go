package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    a := &A{}

    ch := make(chan int, 2)
    go a.Run(true, 1, ch)
    go a.Run(false, 2, ch)


    <-ch
    <-ch
}

func minMoves(nums []int) int {
    l := len(nums)

    if l == 1 {
        return 0
    }

    min := nums[0]
    for _, v := range nums {
        if min > v {
            min = v
        }
    }

    r := 0
    for _, v := range nums {
        r += v-min
    }

    return r
}


type A struct {

}

func (a *A) Run(t bool, i int, c chan int) {
    l := sync.Mutex{}
    l.Lock()
    fmt.Println("lock success")
    defer func() {
        l.Unlock()
        c<-i
    }()

    if t {
        time.Sleep(1 * time.Second) // sleep 5秒
    } else {
        time.Sleep(3 * time.Second) // sleep 5秒
    }

    fmt.Println("end:", i)

}