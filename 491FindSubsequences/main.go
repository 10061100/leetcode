package main

import (
    "fmt"
)

func main()  {
    fmt.Println(findSubsequences([]int{4,6,7,7}))
    fmt.Println(findSubsequences([]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}))
}

func findSubsequences(nums []int) [][]int {
    list := NewArray(len(nums))
    res := make([][]int, 0)

    dfs(0, nums, list, &res)

    return res
}

func dfs(start int, nums []int, list *ArrayList, res *[][]int) {

    pre := -1
    for ; start < len(nums); start++ {
        if nums[start] == pre {
            // 过滤相同的
            continue
        }

        pre = nums[start]
        if list.IsEmpty() || list.LastItem() <= nums[start] {
            list.Add(nums[start])
            if list.Size() > 1 {
                (*res) = append((*res), list.Dump())
            }
            dfs(start+1, nums, list, res)
            list.Del()
        }
    }

}

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

func (n *ArrayList) Size() int {
    return n.End
}