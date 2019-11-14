package main

import (
    "fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func BuildList(array []int) *ListNode {

    if len(array) == 0 {
        return nil
    }

    list := &ListNode{
        Val: array[0],
        Next: nil,
    }

    cur := list

    for i := 1; i < len(array); i++ {
        cur.Next = &ListNode{
            Val: array[i],
            Next: nil,
        }

        cur = cur.Next
    }

    return list
}

func main() {
    // fmt.Println(numComponents(BuildList([]int{0,1,2,3}), []int{0,1,3}))
    // fmt.Println(numComponents(BuildList([]int{0,1,2,3,4}), []int{0,1,3,4}) == 2)
    fmt.Println(numComponents(BuildList([]int{0,1,2,3,4}), []int{0,1,3,4,2})== 1)
    // fmt.Println(numComponents(BuildList([]int{0,1,2,3,4}), []int{0,1,3,2}) == 4)
    // fmt.Println(numComponents(BuildList([]int{0,1,2,3,4}), []int{1,3,4,2}) == 4)
    // fmt.Println(numComponents(BuildList([]int{0,1,2,3,4,5}), []int{5,1,3,4,0}) ==3 )
    // fmt.Println(numComponents(BuildList([]int{0,1,2,3,4}), []int{1}))
}




func numComponents(head *ListNode, G []int) int {
    m := make([]int, 10000)

    for i := 0; i < len(G); i++ {
        m[G[i]] = 1
    }

    // 找到开始的位置
    start := head
    for ; start != nil && m[start.Val] == 0; start = start.Next {}

    if start == nil {
        return 0
    }

    end := start
    res := 0
    t   := 0
    for ; end != nil ; end = end.Next {
        if m[end.Val] == 0 {
            if t != 0 {
                res++
            }
            t = 0
        } else {
            t++
        }
    }

    if t != 0 {
        res++
    }

    return res
}

