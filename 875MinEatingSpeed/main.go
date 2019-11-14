package main

import (
    "fmt"
)

func main()  {
    fmt.Println(minEatingSpeed([]int{3,6,7,11}, 8) == 4)
    fmt.Println(minEatingSpeed([]int{30,11,23,4,20}, 5) == 30)
    fmt.Println(minEatingSpeed([]int{30,11,23,4,20}, 6) == 23)
    fmt.Println(minEatingSpeed([]int{20,20,20,20}, 4) == 20)
    fmt.Println(minEatingSpeed([]int{312884470}, 968709470) == 1) //312884470/ 968709470
}

func minEatingSpeed(piles []int, H int) int {
    sum := 0  // 总和, 最后计算平均值
    right := -1

    for _, v := range piles {
        sum += v
        if v > right {
            right = v
        }
    }
    left := (sum + (H-1))/H

    for ; left < right; {
        mid := (left + right) >> 1
        // fmt.Println(mid)

        if canEat(piles, H, mid) {
            right = mid // mid是有可能吃到的最小的
        } else {
            left = mid + 1
        }

    }

    return right
}



func canEat(piles []int, H int, speed int) bool {
    res := 0
    for _, v := range piles {
        res += (v + speed-1)/speed
    }

    return res <= H
}
