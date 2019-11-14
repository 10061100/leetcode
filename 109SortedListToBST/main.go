package main

import (
    "fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

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

func midTravel(root *TreeNode) {
    if root == nil {
        return
    }

    if root.Left != nil {
        midTravel(root.Left)
    }

    fmt.Print(root.Val, ", ")

    if root.Right != nil {
        midTravel(root.Right)
    }

}

func main() {
    midTravel(sortedListToBST(BuildList([]int{1,2,3})))
    midTravel(sortedListToBST(BuildList([]int{1,2})))
    midTravel(sortedListToBST(BuildList([]int{1})))
    midTravel(sortedListToBST(BuildList([]int{})))
}


func sortedListToBST(head *ListNode) *TreeNode {
    if head == nil {
        return nil
    }

    l := head
    for ; l.Next != nil; l = l.Next {}

    return buildsub(head, l)
}

func NewTree(v int) *TreeNode{
    return &TreeNode{
        Val: v,
        Left:nil,
        Right: nil,
    }
}


func buildsub(start *ListNode, end *ListNode) *TreeNode {
    if start == nil {
        return nil
    }

    var pre *ListNode = nil
    cur := start
    dcur:= start

    for ; dcur != end; {
        pre = cur
        cur = cur.Next
        dcur = dcur.Next
        if dcur == end {
            break
        }
        dcur = dcur.Next
    }

    t := NewTree(cur.Val)

    if pre != nil {
        t.Left = buildsub(start, pre)
    }

    if cur != end {
        t.Right = buildsub(cur.Next, end)
    }

    return t
}