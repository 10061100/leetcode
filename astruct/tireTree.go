package astruct


type TireTreeNode struct {
    isend bool
    dic [26]*TireTreeNode
}

func( n *TireTreeNode) Next(i int) *TireTreeNode {
    if i > 25 || i < 0 {
        return nil
    }

    return n.dic[i]
}


func (n *TireTreeNode) IsEnd() bool {
    return n.isend
}

func NewTireNode() *TireTreeNode {
    return &TireTreeNode{
        isend: false,
    }
}


type TireTree struct {
    head *TireTreeNode
}

func NewTireTree() *TireTree {
    return &TireTree{
        head: NewTireNode(),
    }
}

func( t *TireTree) Add( s string) {

    tmp := t.head
    for _, v := range s {
        n := int(v - 'a')
        if tmp.dic[n] == nil {
            tmp.dic[n] = NewTireNode()
        }

        tmp = tmp.dic[n]
    }

    tmp.isend = true
}

func ( t *TireTree) AddWords( ss []string) {
    for _, v := range ss {
        t.Add(v)
    }
}

func (t *TireTree) Find(ss string) bool {
    if ss == "" {
        return false
    }

    tmp := t.head

    for _, v := range ss {
        tmp = tmp.Next(int(v - 'a'))
        if tmp == nil {
            return true
        }
    }

    return tmp.IsEnd()
}