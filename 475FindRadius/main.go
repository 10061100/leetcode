package main

import (
    "fmt"
    "sort"
)

func main() {
    // fmt.Println(findRadius([]int{1,2,3}, []int{2}) == 1)
    // fmt.Println(findRadius([]int{1,2,3,4}, []int{1,4}) == 1)
    // fmt.Println(findRadius([]int{3}, []int{1,5}) == 2)
    house := []int{282475249,622650073,984943658,144108930,470211272,101027544,457850878,458777923}
    heater:= []int{823564440,115438165,784484492,74243042,114807987,137522503,441282327,16531729,823378840,143542612}
    fmt.Println(findRadius(house, heater))
}

func findRadius(houses []int, heaters []int) int {

    if len(houses) == 0 {
        return 0
    }

    sort.Sort(sort.IntSlice(houses))
    sort.Sort(sort.IntSlice(heaters))

    heaterstart, housestart := 0, 0

    res := -1

    for ; housestart < len(houses); housestart++{
        cur := houses[housestart]
        // 找到第一个 >= 自己的电热器
        for ; heaterstart < len(heaters) && heaters[heaterstart] < cur; heaterstart++ {}

        // 说明最后一个也比自己大
        if heaterstart >= len(heaters) {
            res = max(abs(cur-heaters[len(heaters)-1]), res)
            fmt.Println(housestart, "===>", res)
            continue
        }

        if heaterstart == 0 {
            res = max(abs(cur- heaters[0]), res)
            fmt.Println(housestart, "xxxxxxx>", res)
            continue
        }

        left := abs(cur-heaters[heaterstart-1])
        right:= abs(cur-heaters[heaterstart])

        if left > res && right > res {
            res = left
            if left > right {
                res = right
            }
        }

    }

    return res
}

func max( a, b int) int {
    if a > b {
        return a
    }

    return b
}

func abs(a int) int {
    if a >= 0 {
        return a
    }

    return -a
}
