// LongestWord
package main

func main()  {
    println(longestWord([]string{"w","wo","wor","worl", "world"}))
    println(longestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"}) == "apple")
    println(longestWord([]string{"gbra","jy","pl","zn","gb","j","jyh","jyhm","plr","znicn","p","gbr","zni","znic","aq"}))
}

func longestWord(words []string) string {
    tree := InitDicTrees(words)
    q := NewQueue()

    q.Push(tree)
    ans := ""
    ansl:= 0

    for ; !q.IsEmpty(); {
        t := q.Out()
        if t == nil {
            return ans
        }

        if t.End > 0 || t == tree {

            for _, v := range t.Children {
                if v.End > 0 {
                    q.Push(v)
                }
            }

            if t == tree {
                continue
            }

            w := words[t.End-1]
            if len(w) > ansl {
                ans = w
                ansl= len(w)
            } else if len(w) == ansl && ans > w {
                ans = w
            }

        }

    }

    return ans
}


type DicNode struct {
    Val byte
    Children map[byte]*DicNode
    End int
}


func (d *DicNode) getChild(c byte) *DicNode {
    if v, ok := d.Children[c]; ok {
        return v
    }

    return nil
}


func NewDicNode( v byte, idx int) *DicNode{
    return &DicNode{
        Val:v,
        Children:make(map[byte]*DicNode),
        End:idx,
    }
}


func InitDicTrees(words []string) *DicNode {
    root := NewDicNode('0', 1)
    for i := 0; i < len(words); i++ {
        cur := root
        for _, v := range([]byte(words[i])) {
            child := cur.getChild(v);
            if child == nil {
                cur.Children[v] = NewDicNode(v, 0)
                child = cur.getChild(v)
            }
            cur = child
        }
        cur.End = i+1
    }


    return root
}


type Queue struct {
    C []*DicNode
    N int
}

func NewQueue() *Queue {
    return &Queue{
        C: make([]*DicNode, 0),
        N: 0,
    }
}

func (q *Queue) Push(t *DicNode){
    q.C = append(q.C, t)
    q.N++
}

func (q *Queue) IsEmpty() bool {
    return q.N == 0
}

func (q *Queue) Out() *DicNode {
    if q.IsEmpty() {
        return nil
    }

    q.N--
    t := q.C[0]
    q.C = q.C[1:]
    return t
}