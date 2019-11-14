package main

import (
    "fmt"
)

func main() {
    fmt.Println(brokenCalc(1, 100000))
}


func brokenCalc(X int, Y int) int {
    cache := make(map[int]int)

    return brokenCalcSub(X, Y, cache)
}


func brokenCalcSub(X, Y int, cache map[int]int) int {

    if X >= Y {
        return X - Y
    }

    if v, ok := cache[X]; ok {
        return v
    }

    t := 100000000

    for i := 0; 2 * (X - i) > X; i++ {
        t = min(t , (i+1) + brokenCalcSub(2 * (X - i), Y, cache))
    }


    cache[X] = t

    return t
}


func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}
