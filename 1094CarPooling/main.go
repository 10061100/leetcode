package main

import (
    "fmt"
)

func main() {
    fmt.Println(carPooling([][]int{{2,1,5}, {3,3,7}}, 4) == false)
    fmt.Println(carPooling([][]int{{2,1,5}, {3,3,7}}, 5) == true)
    fmt.Println(carPooling([][]int{{2,1,5}, {3,5,7}}, 3) == true)
    fmt.Println(carPooling([][]int{{3,2,7}, {3,7,9}, {8,3,9}}, 3) == false)
    fmt.Println(carPooling([][]int{{3,2,7}, {3,7,9}, {8,3,9}}, 11) == true)

}

func carPooling(trips [][]int, capacity int) bool {
    start := [101]int{}
    end   := [1001]int{}

    for i := 0; i < len(trips); i++ {
        begin := trips[i][1]
        e   := trips[i][2]
        count := trips[i][0]

        start[begin] += count
        end[e] -= count
    }

    t := 0
    for i := 0; i < 101; i++ {
        t += end[i]
        t += start[i]

        if t > capacity {
            return false
        }
    }

    return true
}