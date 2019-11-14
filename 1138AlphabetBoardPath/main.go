package main

import (
    "fmt"
)

func main() {

    fmt.Println((alphabetBoardPath("leetz"))) // "DDR!UURRR!!DDD!"
    // fmt.Println(string(dfs(0, 4, 'l', graph)))
    // fmt.Println(string(dfs(0, 1, 5,0 , graph)))
}

var graph = []string{"abcde", "fghij", "klmno", "pqrst", "uvwxy", "z"}

func alphabetBoardPath(target string) string {

    curx, cury := 0, 0

    mx := make(map[uint8]int)
    my := make(map[uint8]int)

    for x, str := range graph {
        for y := range str {
            mx[str[y]] = x
            my[str[y]] = y
        }
    }

    res := make([]byte, 0)
    for x := range target {
        desx := mx[target[x]]
        desy := my[target[x]]

        res = append(res, dfs(curx, cury, desx, desy, graph)...)

        res = append(res, '!')
        curx, cury = desx, desy
    }

    return string(res)
}

func dfs(curx, cury int, desx, desy int, graph []string) (res []byte) {
    res = make([]byte, 0)

    if curx == desx && cury == desy {
        return res
    }

    if graph[desx][desy] == 'z' {
        t := dfs(curx, cury, 4, 0, graph) // 先去u,再去z
        res = append(t, 'D')
        return res
    }

    if desx > curx {
        res = append(res, repeat(desx-curx, 'D')...)
    } else if curx > desx {
        res = append(res, repeat(curx-desx, 'U')...)
    }

    if desy > cury {
        res = append(res, repeat(desy-cury, 'R')...)
    } else if cury > desy {
        res = append(res, repeat(cury-desy, 'L')...)
    }

    return res
}


func repeat(n int, b byte) []byte {
    res := make([]byte, 0)

    for i := 0; i < n; i++ {
        res = append(res, b)
    }

    return res
}