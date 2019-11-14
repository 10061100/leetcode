package main

import (
    "fmt"
    "strconv"
)

func main() {
    s := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

    fmt.Println(groupAnagrams(s))
}

func groupAnagrams(strs []string) [][]string {
    t := NewTool()

    tmp := make(map[string][]string)

    for _, v := range strs {
        hash := t.Init(v).HashCode()
        if _, ok := tmp[hash]; !ok {
            tmp[hash] = make([]string, 0)
        }

        tmp[hash] = append(tmp[hash], v)
    }

    res := make([][]string, 0)
    for _, v := range tmp {
        res = append(res, v)
    }

    return res
}


func NewTool() *Tool {
    return &Tool{}
}

type Tool struct {
    index [26]int
}

func (t *Tool) Init(s string) *Tool {
    t.ToEmpty()

    for _, v := range []byte(s) {
        t.index[v - 'a']++
    }

    return t
}


func ( t *Tool) ToEmpty() {
    for k :=  range t.index {
        t.index[k] = 0
    }
}

func (t *Tool) HashCode() string {
    s := ""
    for k, v := range t.index {
        if v == 0 {
            continue
        }
        s += string('a' + k) + strconv.Itoa(v)
    }

    return s
}
