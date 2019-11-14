package main

import (
    "fmt"
    "sort"
)

func main() {
    // fmt.Println(findItinerary([][]string{{"JFK", "MUC"}, {"MUC", "ABC"}}))
    // fmt.Println(findItinerary([][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}))
    fmt.Println(findItinerary([][]string{{"JFK","SFO"},{"JFK","ATL"},{"SFO","ATL"},{"ATL","JFK"},{"ATL","SFO"}}))
}

func findItinerary(tickets [][]string) []string {
    m := NewMap(tickets)

    if m.Len() == 0 {
        return nil
    }
    res := make([]string, m.Len()) // 包括了JFK
    vis := make(map[string]bool) // 是否已经访问了

    dfs(vis, m.m["JFK"], m.Len(), 0, res)

    return res
}

func dfs(vis map[string]bool, start *MapNode, total, cur int, res []string ) bool{
    if start == nil {
        return false
    }

    if cur >= total {
        return true
    }


    res[cur] = start.Val

    if cur == total- 1 {
        return true
    }

    for _, v := range start.Next {
        key := start.Val + v.Val
        if vis[key] == true {
            continue
        }
        vis[key] = true
        if dfs(vis, v, total, cur+1, res) {
            return true
        }
        vis[key] = false
    }

    return false
}

type Map struct{
    m map[string]*MapNode
}

func (m *Map) Len() int {
    return len(m.m)
}


func NewMap(con [][]string) *Map {
    m := &Map{m:make(map[string]*MapNode)}

    for _, v := range con {
        start := v[0]
        end   := v[1]

        if _, ok := m.m[start]; !ok {
            m.m[start] = NewNode(start)
        }

        _, ok := m.m[end]
        if !ok {
            m.m[end] = NewNode(end)
        }

        enode := m.m[end]

        m.m[start].Next = append(m.m[start].Next, enode)
    }

    for k := range m.m {
        m.m[k].SortNext()
    }

    return m
}


type MapNode struct {
    Val string
    Next []*MapNode
}

func NewNode(v string) *MapNode{
    return &MapNode{
        Val: v,
        Next: make([]*MapNode, 0),
    }
}

func (m *MapNode) SortNext() {
    sort.Sort(MapArray(m.Next))
}

type MapArray []*MapNode

func (m MapArray) Len() int {
    return len(m)
}

func (m MapArray) Less(i ,j int) bool {
    return m[i].Val < m[j].Val
}

func (m MapArray) Swap(i, j int) {
    m[i], m[j] = m[j], m[i]
}