package main

import (
    "fmt"
)

func main() {
    fmt.Println(numRollsToTarget(1,6,3) == 1)
    fmt.Println(numRollsToTarget(2,6,7) == 6)
}



const F = 1000000007

func numRollsToTarget(d int, f int, target int) int {

    cache := make([]map[int]int, 31)

    for i := 0; i <= 30; i++ {
        cache[i] = make(map[int]int)
    }

    return subProblem(d, f, target, cache)

}


func subProblem(d int, f int, target int, cache []map[int]int ) int {

    if _, ok := cache[d][target]; ok {
        return cache[d][target]
    }

    if d == 1 {
        if target <= f {
            return 1
        }

        return 0
    }

    s := 0
    for i := 1; i <= f; i++ {
        s += subProblem(d-1, f, target-i, cache)
        s = s % F
    }
    cache[d][target] = s

    return cache[d][target]
}