package main

import (
    "fmt"
)

func main() {
    fmt.Println(allPathsSourceTarget([][]int{{1,2},{3},{3},{}}))
}

func allPathsSourceTarget(graph [][]int) [][]int {
    res := make([][]int, 0)
    n := len(graph)
    list := NewArray(n)
    list.Add(0)

    dfs(graph, 0, n-1, list, &res)

    return res
}


func dfs(graph [][]int, node int, des int, list *ArrayList, res *[][]int) {

    if node == des {
        *res = append(*res, list.Dump())
        return
    }

    for _, next := range graph[node] {
        list.Add(next)
        dfs(graph, next, des, list, res)
        list.Del()
    }
}




type ArrayList struct {
    End int
    Values []int
    Len int
}


func NewArray(n int) *ArrayList{
    if n < 0 {
        n = 100
    }

    return &ArrayList{
        End: 0,
        Values: make([]int, n),
        Len: n,
    }
}


func (n *ArrayList) Add(v int) {
    if n.End < n.Len {
        n.Values[n.End] = v
    } else {
        n.Len++
        n.Values = append(n.Values, v)
    }

    n.End++
}

func (n *ArrayList) Del() {
    if n.End > 0 {
        n.End--
    }
}

func (n *ArrayList) Dump() []int {
    t := make([]int, n.End)

    copy(t, n.Values[0:n.End])

    return t
}