package main

import (
    "fmt"
)

func main() {
    PrintList(reorderList(BuildList([]int{1,2,3})))
}

type ListNode struct {
    Val int
    Next *ListNode
}


func PrintList(list *ListNode) {
    fmt.Print("[")
    defer fmt.Println("]")
    i := 0
    for cur := list; i <= 10 && cur != nil; cur, i = cur.Next, i+1 {
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


func reorderList(head *ListNode)  *ListNode{
    if head == nil || head.Next == nil {
        return nil
    }

    mid := findMidList(head)
    t := mid
    mid = mid.Next
    t.Next = nil

    mid = reverseList(mid)
    cur1, cur2 := head.Next, mid
    curres := head
    curres.Next = nil

    for i := 1; cur1 != nil && cur2 != nil; i++ {
        if i % 2 == 1 {
            // 选取list2的值
            t := cur2
            cur2 = cur2.Next
            curres.Next = t
            t.Next = nil
            curres = t
        } else {
            t := cur1
            cur1 = cur1.Next
            curres.Next = t
            t.Next = nil
            curres = t
        }
    }

    if cur2 != nil {
        curres.Next = cur2
    }

    if cur1 != nil {
        curres.Next = cur1
    }

    return head
}


func findMidList(head *ListNode) *ListNode {
    l1, l2 := head, head.Next

    for ; l2 != nil; {
        l1 = l1.Next
        l2 = l2.Next
        if l2 != nil {
            l2 = l2.Next
        }
    }

    return l1
}


func reverseList(list *ListNode) *ListNode {
    if list == nil {
        return list
    }

    pre , cur := list, list.Next
    pre.Next = nil

    for ; cur != nil; {
        t := cur
        cur = cur.Next
        t.Next = pre
        pre = t
    }

    return pre
}