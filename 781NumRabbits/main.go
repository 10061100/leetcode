package main

import (
    "fmt"
    "sort"
)

func main()  {
    fmt.Println(numRabbits([]int{1,1,1}) == 4)
    fmt.Println(numRabbits([]int{2}) == 3)
    fmt.Println(numRabbits([]int{2,2,2}) == 3)
    fmt.Println(numRabbits([]int{10,10,10}) == 11)
}

func numRabbits(answers []int) int {
    if len(answers) == 0 {
        return 0
    }

    sort.Sort(sort.IntSlice(answers))

    t := 0
    c := 1
    v := answers[0]
    for i := 1; i < len(answers); i++ {
        if answers[i] == v {
            c++
        } else {
            t += (c + v)/(v+1) * (v+1)
            // 否则c == 1
            c = 1
            v = answers[i]
        }
    }

    t += (c + v)/(v+1) * (v+1)

    return t
}
