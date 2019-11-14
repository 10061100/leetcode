package main

import (
    "fmt"
)

func main() {
    fmt.Println(maxSatisfied([]int{1,0,1,2,1,1,7,5}, []int{1,1,1,1,1,1,1,1}, 2))
    fmt.Println(maxSatisfied([]int{1,0,1,2,1,1,7,5}, []int{0,1,0,1,0,1,0,1}, 3))
    fmt.Println(maxSatisfied([]int{1,0,1,2,1,1,7,5}, []int{1,1,1,1,1,1,1,1}, 1))
    fmt.Println(maxSatisfied([]int{4,10,10}, []int{1,1,0}, 2))
    fmt.Println(maxSatisfied([]int{2,4,1,4,1}, []int{1,0,1,0,1}, 2))
}

func maxSatisfied(customers []int, grumpy []int, X int) int {
    l := len(customers)

    if l == 0 {
        return 0
    }

    res, max := 0, 0
    left, right := 0, 0

    for i := 0; i < l; i++ {
        if grumpy[i] == 0 {
            res += customers[i]
        }
    }

    for ; ;  {
        // 找到左指针的位置
        for ; left < l && grumpy[left] == 0; left++{}

        // 首先找到符合要求右指针
        for ; right < l ; right++ {
            if grumpy[right] == 1 && right >= left + X {
                break
            }

            if grumpy[right] == 1 {
                res += customers[right]
            }
        }

        if max < res {
            max = res
        }

        // fmt.Println(res)
        if left >= l {
            break
        }

        // 这个时候,开始启动下一个周期, left
        res -= customers[left]
        left++
    }

    return max
 }

