package main

import (
    "fmt"
)

func main() {

    fmt.Println(shortestAlternatingPaths(3, [][]int{{0,1},{1,2}}, [][]int{{1,1}}))
    fmt.Println(shortestAlternatingPaths(3, [][]int{{0,1},{1,2}}, [][]int{}))
    fmt.Println(shortestAlternatingPaths(3, [][]int{{0,1}}, [][]int{{2,1}}))
    fmt.Println(shortestAlternatingPaths(3, [][]int{{1,0}}, [][]int{{2,1}}))
    fmt.Println(shortestAlternatingPaths(3, [][]int{{0,1}}, [][]int{{1,2}}))
    fmt.Println(shortestAlternatingPaths(8,
        [][]int{{7,2},{4,1},{5,3},{0,3},{5,6},{0,4},{6,2},{4,7},{3,0},{2,6},{1,2},{2,2},{5,7},{7,1},{3,6},{3,2},{6,6},{1,4},{4,4},{7,6},{6,1},{5,1},{7,7},{6,4}},
        [][]int{{0,7},{3,4},{1,0},{0,0},{4,6},{1,7},{3,7}}))
}

func shortestAlternatingPaths(n int, red_edges [][]int, blue_edges [][]int) []int {

    // 尾节点是k, 且最后一条指向k点的边是红色或蓝色的长度
    red     := make([]int, n)
    blue    := make([]int, n)

    g := make([][]int, n)
    for i := 0; i < n; i++ {
        g[i] = make([]int, n)
        red[i] = -1
        blue[i] = -1
    }

    blue[0] ,red[0] = 0, 0
    for _, v := range red_edges {
        g[v[0]][v[1]] = g[v[0]][v[1]] | 1
    }

    for _, v := range blue_edges {
        g[v[0]][v[1]] = g[v[0]][v[1]] | 2
    }

    // fmt.Printf("%v\n", g)
    // 每次走一步
    pre := []int{0}
    for ; len(pre) != 0 ; {
        cur := make([]int, 0)
        // fmt.Println(pre)
        for _, start := range pre {
            // fmt.Println("START:",start)
            for des, edge := range g[start] {
                // 是红边
                add := -1
                // 这块有个case, 必须要以当前的镜像为准进行大小判断, 否则就会出现接力的情况
                if edge & 1 != 0 && blue[start] != -1 && (red[des] == -1 || (blue[start] + 1 < red[des])){
                    red[des] = blue[start] + 1
                    add = des
                }

                // 是蓝边
                if edge & 2 != 0 && red[start] != -1 && (blue[des] == -1 || (red[start] + 1 < blue[des])){
                    blue[des] = red[start] + 1
                    add = des
                }

                if add != -1 {
                    cur = append(cur, add)
                }
            }

            // fmt.Println("=======")
            // fmt.Println(blue)
            // fmt.Println(red)
        }
        pre = cur
    }

    // fmt.Println(blue)
    // fmt.Println(red)
    res := make([]int, n)

    for i := 0; i < n; i++ {
        if blue[i] != -1 && red[i] != -1 {
            res[i] = min(blue[i], red[i])
            continue
        }

        res[i] = blue[i]
        if blue[i] == -1 {
            res[i] = red[i]
        }
    }

    return res
}

func min(a, b int) int {
    if a > b {
        return b
    }

    return a
}
