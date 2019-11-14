package main

import (
    "fmt"
    "sort"
)

func main()  {
    tm := Constructor()

    tm.Set("123", "high", 10)
    tm.Set("123", "low", 20)
    // tm.Set("123", "3", 30)

    // fmt.Println(tm.Get("123",5) == "")
    // fmt.Println(tm.Get("123",10) == "high")
    fmt.Println(tm.Get("123",15))
    // fmt.Println(tm.Get("123",20) == "2")
    // fmt.Println(tm.Get("123",31) == "3")
    // fmt.Println(tm.Get("123",30) == "3")
    // fmt.Println(tm.Get("123",9) == "1")
    // fmt.Println(tm.Get("123",1) == "1")
    // fmt.Println(tm.Get("123",0) == "1")
}

type Pair struct {
    Val string
    Time int
}

type TimeMap struct {
    Field map[string][]*Pair
}


/** Initialize your data structure here. */
func Constructor() TimeMap {
    return TimeMap{
        Field: make(map[string][]*Pair),
    }
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
    if _, ok := this.Field[key]; !ok {
        this.Field[key] = make([]*Pair, 0)
    }

    this.Field[key] = append(this.Field[key], &Pair{Val: value, Time: timestamp})
}


func (this *TimeMap) Get(key string, timestamp int) string {
    if _, ok := this.Field[key]; !ok {
        return ""
    }
    t := this.Field[key]

    k := sort.Search(len(t), func(i int) bool {
        if t[i].Time >= timestamp {
            return true
        }

        return false
    })

    if k >= len(t) {
        return t[k-1].Val
    }

    fmt.Println(t[k])
    if t[k].Time > timestamp {
        return ""
    }

    return t[k].Val
}

