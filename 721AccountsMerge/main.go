package main

import (
    "fmt"
    "sort"
)

func main() {
    fmt.Println(accountsMerge([][]string{{"John", "johnsmith@mail.com", "john00@mail.com"},
        {"John", "johnnybravo@mail.com"},
        {"John", "johnsmith@mail.com", "john_newyork@mail.com"},
        {"Mary", "mary@mail.com"},}))
}


func accountsMerge(accounts [][]string) [][]string {

    email2name := make(map[string]string)

    set := NewRebinSet()

    for _, v := range accounts {
        name := v[0]
        first := ""
        for i := 1; i < len(v); i++ {
            email2name[v[i]] = name
            // 如果不存在
            if !set.Contain(v[i]) {
                set.Add(v[i])
            }

            if i == 1 {
                first = v[i]
                continue
            }

            set.Set(first, v[i])
        }
    }

    cache := make(map[string][]string)
    for email := range email2name {
        key := set.GetRoot(email)
        if _, ok := cache[key]; !ok {
            cache[key] = make([]string, 0)
        }

        cache[key] = append(cache[key], email)
    }

    res := make([][]string, 0)
    for k, v := range cache {
        name := email2name[k]
        sort.Sort(sort.StringSlice(v))
        res = append(res, append([]string{name}, v...))
    }

    return res
}



type RebinSet struct {
    M map[string]string
}


func NewRebinSet() *RebinSet {
    return &RebinSet{
        M: make(map[string]string),
    }
}


func (r *RebinSet) GetRoot(s string) string {

    if _, ok := r.M[s]; !ok {
        return s // 没有的话,  就是他自己
    }

    for ; ; {
        if p, ok := r.M[s]; !ok || p == "" {
            return s
        } else {
            s = p
        }
    }
}

func (r *RebinSet) Set(parent, s string) {

    if !r.Contain(s) {
        r.Add(s)
    }

    p := r.GetRoot(parent)

    sp := r.GetRoot(s)

    if p > sp {
        r.M[sp] = p
    } else if p < sp {
        r.M[p] = sp
    }

}

func (r *RebinSet ) Contain(s string) bool {
    if _, ok := r.M[s]; !ok {
        return false
    }

    return true
}

func (r *RebinSet) Add(s string) {
    r.M[s] = "" // 自己是自己的祖先
}

func (r *RebinSet) OutPut() {
    for  k, v := range r.M {
        println(k, "=>", v)
    }
}

