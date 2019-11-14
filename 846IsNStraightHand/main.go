package main

import (
    "fmt"
    "sort"
)

func main()  {
    fmt.Println(isNStraightHand([]int{1,2,3,6,2,3,4,7,8}, 3) == true)
    fmt.Println(isNStraightHand([]int{1,2,3,4,5,6,7,8,9}, 2) == false)
    fmt.Println(isNStraightHand([]int{1,2,3,4,5,6,7,8,9}, 1) == true)
    fmt.Println(isNStraightHand([]int{1,2,3,7,8,9}, 3) == true)
    fmt.Println(isNStraightHand([]int{1,2,3,7,9,9}, 3) == false)
}

func isNStraightHand(hand []int, W int) bool {

    lh := len(hand)

    if lh % W != 0 {
        return false
    }

    sort.Sort(sort.IntSlice(hand))

    p := make([]int, W)
    p[0] = 0
    visited := make([]bool, lh)

    for i := 0; i < W; i++ {
        p[i] = 0
    }

    for ; p[0] < lh; {
        pre := -1
        i := 0
        for ; i < W; i++ {
            for ; p[i] < lh ; p[i]++ {
                if visited[p[i]] {
                    continue
                }

                if pre == -1 || hand[p[i]] == pre + 1 {
                    pre = hand[p[i]]
                    visited[p[i]] = true
                    break
                }
            }

            if i == 0 && p[i] >= lh {
                break
            }

            if i !=0 && p[i] >= lh {
                // fmt.Println(i, ",", p[i], ",", lh)
                return false
            }
        }
    }

    return true
}
