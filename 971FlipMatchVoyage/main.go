package main


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


func flipMatchVoyage(root *TreeNode, voyage []int) []int {
    return nil
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
