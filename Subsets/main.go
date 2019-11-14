package main

import (
    "fmt"
)

func main()  {
    fmt.Println(subsets([]int{1,2,3}))
}

func subsets(nums []int) [][]int {

    res := make([][]int, 0)
    if len(nums) == 0 {
        return res
    }

    repeat(NewArray(), &res, nums, 0)

    // fmt.Println(res)

    return  res
}


func repeat( t *Array, res *[][]int, nums []int, start int) {
    l := len(nums)

    if start >= l {
        *res = append(*res, t.Dump())
        // fmt.Println(res)
        return
    }

    // t.Add(nums[start])
    // repeat(t, res, nums, start+1)
    // t.Del()
    // repeat(t, res, nums, start+1)

    *res = append(*res, t.Dump())
    for i := start; i < l; i++ {
        // repeat(t, res, nums, i+1)
        t.Add(nums[i])
        repeat(t, res, nums, i+1)
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