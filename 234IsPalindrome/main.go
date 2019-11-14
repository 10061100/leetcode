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
    fmt.Println(isPalindrome(BuildList([]int{1,1})) == true)
    fmt.Println(isPalindrome(BuildList([]int{1})) == true)
    fmt.Println(isPalindrome(BuildList([]int{1,2,1})) == true)
    fmt.Println(isPalindrome(BuildList([]int{1,2,3})) == false)
    fmt.Println(isPalindrome(BuildList([]int{1,2,2,1})) == true)
}

func isPalindrome(head *ListNode) bool {
    if head == nil {
        return true
    }

    l := 0

    for t := head; t != nil; t, l = t.Next, l+1 {}

    pre :=  l/2

    // fmt.Println(pre)
    // 顺序逆转前面pre个的数据链表
    var newlist *ListNode = nil
    var tail = head
    for ; pre >= 1; pre-- {
        next := tail.Next
        tail.Next = newlist
        newlist = tail
        tail = next
    }

    if l % 2 == 1 {
        tail = tail.Next
    }

    n1, n2 := newlist, tail
    for ; n1 != nil && n2 != nil; {
        if n1.Val == n2.Val {
            n1 = n1.Next
            n2 = n2.Next
            continue
        }

        return false
    }

    // fmt.Println(n1)
    // fmt.Println(n2)
    if n1 != nil || n2 != nil {
        return false
    }

    return true
}
