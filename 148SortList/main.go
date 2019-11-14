package main

import (
    "fmt"
)

func main() {
    fmt.Println(sortList(BuildList([]int{4,2,1,3})))
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

type ListNode struct {
    Val int
    Next *ListNode
}

func sortList(head *ListNode) *ListNode {
    h, t := quickSort(head)
    t.Next = nil

    return h
}



func quickSort(from *ListNode) (head, end *ListNode) {
    if from == nil {
        return nil, nil
    }

    cur := from
    from = from.Next
    cur.Next = nil
    var big, curbig *ListNode
    var small, cursmall *ListNode
    for ; from != nil; from = from.Next {
        if from.Val >= cur.Val {
            if big == nil {
                big = from
                curbig = big
            } else {
                curbig.Next = from
                curbig = from
            }
        } else {
            if small == nil {
                small = from
                cursmall = small
            } else {
                cursmall.Next = from
                cursmall = from
            }
        }
    }

    from = cur
    to   := cur
    to.Next = nil
    if small != nil {
        smallhead, smalltail := quickSort(small)
        from = smallhead
        smalltail.Next = cur
    }

    if big != nil {
        bighead, bigtail := quickSort(big)
        to.Next = bighead
        bigtail.Next = nil
        to = bigtail
    }

    return from, to
}