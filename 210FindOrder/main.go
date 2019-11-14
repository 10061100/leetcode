package main

import (
    "fmt"
)

func main() {
    // fmt.Println(findOrder(4, [][]int{{1,0},{2,0},{3,1},{3,2}}))
    // fmt.Println(findOrder(2, [][]int{{1,0}}))
    // fmt.Println(findOrder(2, [][]int{{0,1}}))
    fmt.Println(findOrder(3, [][]int{{1,0}, {1,2}, {0,1}}))
}

func findOrder(numCourses int, prerequisites [][]int) []int {
    degree := make([]int, numCourses)
    link := make([][]int, numCourses)
    for i := 0; i < numCourses; i++ {
        link[i] = make([]int, 0)
    }

    for _, v := range prerequisites {
        end := v[0]
        start := v[1]

        degree[end]++
        link[start] = append(link[start], end)
    }

    res := make([]int, 0)
    visted := make([]int, numCourses)
    for i := 0; i < numCourses; i++ {
        if degree[i] == 0 && visted[i] != 1 {
            if t := dfs(i, link, degree, visted, &res); t == false {
                return make([]int, 0)
            }
        }
    }

    for i := 0; i < numCourses; i++ {
        if degree[i] > 0 {
            return make([]int, 0)
        }
    }

    return res
}

func dfs(start int, link [][]int, degree []int, visited []int, res *[]int) bool {

    if visited[start] == 2 {
        return false
    }

    if degree[start] > 0 {
        return true
    }

    *res = append(*res, start)
    // 节点的前置节点已经完成
    visited[start] = 2
    for i := 0; i < len(link[start]); i++ {
        degree[link[start][i]]--
        if t := dfs(link[start][i], link, degree, visited, res); t == false {
            return t
        }
    }
    visited[start] = 1

    return true
}