package main

import (
    "fmt"
    "sort"
)

func main() {
    fmt.Println(combinationSum([]int{2,3,5}, 8))
}

func combinationSum(candidates []int, target int) [][]int {
    sort.Ints(candidates)

    res := make([][]int, 0)

    repeat(NewArray(), &res, candidates, 0, target)

    return res
}



func repeat( t *Array, res *[][]int, nums []int, start int, limit int) {
    l := len(nums)

    if limit == 0 {
        *res = append(*res, t.Dump())
        return
    }

    if start >= l {
        return
    }

    // t.Add(nums[start])
    // repeat(t, res, nums, start+1)
    // t.Del()
    // repeat(t, res, nums, start+1)

    // *res = append(*res, t.Dump())
    for i := start; i < l; i++ {
        // repeat(t, res, nums, i+1)
        if nums[i] > limit {
            continue
        }

        t.Add(nums[i])
        repeat(t, res, nums, i, limit-nums[i])
        t.Del()
    }
}


type Array struct  {
    store []int
    cur   int
    length int
}


func NewArray() *Array{
    return &Array{
        store: make([]int, 1),
        cur  : 0,
        length: 1,
    }
}

func (t *Array) ToEmpty() {
    t.cur = 0
}

func (t *Array) Add(v int) {
    if t.cur >= t.length {
        t.store = append(t.store, v)
        t.length++
    } else {
        t.store[t.cur] = v
    }

    t.cur++
}

func (t *Array) Del() {
    if t.cur > 0 {
        t.cur--
    }
}

func (t *Array) Dump() []int {
    res := make([]int, t.cur)

    copy(res, t.store[:t.cur])

    return res
}