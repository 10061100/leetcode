package main

import (
    "fmt"
)

func main() {
    fmt.Println(leastBricks([][]int{{1,2,2,1},{3,1,2},{1,3,2},{2,4},{3,1,2},{1,3,1,1}}))
}


func leastBricks(wall [][]int) int {
    m := make(map[int]int)

    res := 0
    for i := 0; i < len(wall); i++ {
        t := 0
        for j := 0; j < len(wall[i]) - 1; j++ {
            t += wall[i][j]
            if _, ok := m[t]; !ok {
                m[t] = 0
            }

            m[t]++

            if m[t] > res {
                res = m[t]
            }

        }
    }

    return len(wall) - res
}



