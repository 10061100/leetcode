// SelfDividingNumbers
package main

import "fmt"

func main() {
    fmt.Println(selfDividingNumbers(1,22))

}


func selfDividingNumbers(left int, right int) []int {
    res := make([]int, 0)
    for i := left ; i <= right; i++ {
        if isSelfDivid(i) {
            res = append(res, i)
        }
    }

    return res
}


func isSelfDivid(n int) bool {
    t := n
    for t > 0 {
        l := t % 10
        if l == 0 || n % l != 0 {
            return false
        }

        t /= 10
    }

    return true
}