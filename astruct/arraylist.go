package astruct


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
    copy(t, n.Values[:n.End])

    return t
}

func (n *ArrayList) IsEmpty() bool {

    return n.End == 0
}


func (n *ArrayList) LastItem() int {
    if n.IsEmpty() {
        return -1
    }

    return n.Values[n.End-1]
}