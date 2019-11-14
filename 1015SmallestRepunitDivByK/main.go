package main

import (
    "fmt"
)

func main()  {
    fmt.Println(smallestRepunitDivByK(3) == 3)
    fmt.Println(smallestRepunitDivByK(4) == -1)
    fmt.Println(smallestRepunitDivByK(2) == -1)
    fmt.Println(smallestRepunitDivByK(5) == -1)
    fmt.Println(smallestRepunitDivByK(111) ==3)
    fmt.Println(smallestRepunitDivByK(17) == 16)
    fmt.Println(smallestRepunitDivByK(19) == 18)
    fmt.Println(smallestRepunitDivByK(13) == 6)
}


func smallestRepunitDivByK(K int) int {

    t := 1
    exist := make([]bool, 100001)
    res := 1

    for ; ; {

        if t % K == 0 {
            break
        }

        if exist[t%K] == true {
            return -1
        }

        exist[t%K] = true

        t = (t % K) * 10 + 1
        res++
    }

    return res
}
