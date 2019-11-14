package astruct

import (
    "fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}


func PrintList(list *ListNode) {
    fmt.Print("[")
    defer fmt.Println("]")

    for cur := list; cur != nil; cur = cur.Next {
        fmt.Print(cur.Val, "\t,")
    }

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