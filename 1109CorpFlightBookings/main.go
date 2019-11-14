package main

import (
    "fmt"
)

func main() {
    fmt.Println(corpFlightBookings([][]int{{1,2,10}, {2,3,20}, {2,5,25}}, 5))
}


func corpFlightBookings(bookings [][]int, n int) []int {
    hk := make([]int, n)

    for _, v := range bookings {
        start := v[0]-1
        end   := v[1]-1
        count := v[2]

        hk[start] += count
        if end < n - 1 {
            hk[end+1] -= count
        }
    }

    for i := 1; i < n; i++ {
        hk[i] = hk[i-1] + hk[i]
    }

    return hk
}


