package main




func findCircleNum(M [][]int) int {
    l := len(M)

    set := NewContinusRebinSet(l)

    for i := 0; i < l; i++ {
        for j := 0; j < l; j++ {
            if M[i][j] == 1 {
                set.Set(i, j)
            }
        }
    }

    c := 0

    for _, v := range set.M {
        if v == -1 {
            c++
        }
    }

    return c

}


type ContinusRebinSet struct {
    M []int
    Len int
}


func NewContinusRebinSet(n int) *ContinusRebinSet {

    if n <= 0 {
        n = 100
    }


    c := &ContinusRebinSet{
        M: make([]int, n),
        Len: n,
    }

    for i := 0; i < n; i++ {
        c.M[i] = -1
    }

    return c
}

func (c *ContinusRebinSet) GetRoot(k int) int {

    if c.M[k] == -1 {
        return k
    }

    for ; c.M[k] != -1; k = c.M[k] {}

    return k
}


func (c *ContinusRebinSet) Set(parent int, son int) {

    proot := c.GetRoot(parent)
    sroot := c.GetRoot(son)

    // TODO 如果并查集发现有共同的父亲, 就优化son的节点
    // if proot == sroot {
    //     c.M[sroot] =
    // }

    if proot != sroot {
        if proot > sroot {
            c.M[proot] = sroot
        } else {
            c.M[sroot] = proot
        }
    }

}

