package main

import (
    "fmt"
)

func main() {
    fmt.Println(lexicalOrder(199))
}


func lexicalOrder(n int) []int {
    list := NewArray(n)
    for i := 1; i <= 9;  i++ {
        if i > n {
            break
        }

        list.Add(i)
        add(n, i, list)
    }

    return list.Values
}

func add(n int, pre int, list *ArrayList) {

    pre = pre * 10

    // list.Add(pre)
    for i := 0; i <= 9; i++ {
        if pre > n {
            return
        }
        list.Add(pre)
        add(n, pre, list)
        pre++
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
