package main

import (
    "fmt"
)

func main() {
    // fmt.Println(replaceWords([]string{"cat", "baty", "rat"}, "the cattle was rattled by the battery"))
    tree := NewTireTree()
    tree.AddWords([]string{"a", "aaa"})
    fmt.Println(tree.Find("aaaa"))
}


func replaceWords(dict []string, sentence string) string {
    if len(dict) == 0 {
        return sentence
    }

    if sentence == "" {
        return sentence
    }

    tree := NewTireTree()
    tree.AddWords(dict)

    start := 0
    res := make([]byte, 0)
    bs  := []byte(sentence)
    for i := 0; i < len(bs); i++ {
        if (i < len(bs) - 1 && bs[i+1] != ' ') {
            // fmt.Println(i)
            continue
        }

        // fmt.Println("======", string(bs[start:i+1]))
        // fmt.Println(tree.Find(string(bs[start:i+1])))
        if ok, v := tree.Find(string(bs[start:i+1])); ok {
            res = append(res, []byte(v)...)
        } else {
            res = append(res, bs[start:i+1]...)
        }

        if i < len(bs) - 1 {
            res = append(res, bs[i+1])
        }
        start = i+2
    }

    return string(res)
}


type TireTreeNode struct {
    isend bool
    dic [26]*TireTreeNode
    endWord string
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

    if s == "" {
        return
    }

    tmp := t.head
    for _, v := range s {
        n := int(v - 'a')
        if tmp.dic[n] == nil {
            tmp.dic[n] = NewTireNode()
        }

        tmp = tmp.dic[n]
    }

    tmp.isend = true
    tmp.endWord = s
}

func ( t *TireTree) AddWords( ss []string) {
    for _, v := range ss {
        t.Add(v)
    }
}

func (t *TireTree) Find(ss string) (bool, string) {
    if ss == "" {
        return false, ""
    }

    tmp := t.head

    for _ , v := range ss {
        next := tmp.Next(int(v - 'a'))
        if next == nil {
            return tmp.IsEnd(), tmp.endWord
        }
        if next.IsEnd() {
            return next.IsEnd(), next.endWord
        }

        tmp = next
    }

    return tmp.IsEnd(), tmp.endWord
}
