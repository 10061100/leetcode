
// MergeTwoLists
package main

func main()  {
    result := mergeTwoLists(getL1(), getL2())
    // result := getL1()
    printList(result)
}

type ListNode struct {
    Val int
    Next *ListNode
}

func getL1() *ListNode {
    l1 := &ListNode{Val:1, Next:nil}
    cur := l1
    cur.Next = &ListNode{Val:2, Next:nil}
    cur = cur.Next
    cur.Next = &ListNode{Val:4, Next:nil}
    return l1
}

func getL2() *ListNode {
    l2 := &ListNode{Val:1, Next:nil}
    cur := l2
    cur.Next = &ListNode{Val:3, Next:nil}
    cur = cur.Next
    cur.Next = &ListNode{Val:4, Next:nil}
    return l2
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if nil == l1 {
        return l2
    }

    if nil == l2 {
        return l1
    }


    header := l2
    if l1.Val < l2.Val {
        header = l1
        l1 = l1.Next
    } else {
        l2 = l2.Next
    }
    cur := header

    for ; l1 != nil && l2 != nil;  {
        if l1.Val <= l2.Val {
            cur.Next = l1
            l1 = l1.Next
        } else {
            cur.Next = l2
            l2 = l2.Next
        }
        cur = cur.Next
    }

    if l1 != nil {
        cur.Next = l1
    } else if l2 != nil {
        cur.Next = l2
    }

    return header
}


func printList(result *ListNode) {
    for ; nil != result; {
        print(result.Val)
        print(",")
        result = result.Next
    }
}
