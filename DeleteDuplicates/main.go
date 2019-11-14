// DeleteDuplicates
package main

func main() {
    PrintNode(deleteDuplicates(getNode()))
}

func getNode() *ListNode {
    list := &ListNode{1,nil}

    l := list
    l.Next = &ListNode{1,nil}

    l = l.Next
    l.Next = &ListNode{1,nil}

    l = l.Next
    l.Next = &ListNode{1,nil}
    l = l.Next
    l.Next = &ListNode{1,nil}
    l = l.Next
    l.Next = &ListNode{1,nil}
    l = l.Next
    l.Next = &ListNode{2,nil}

    l = l.Next
    l.Next = &ListNode{3,nil}

    return list
}

type ListNode struct {
    Val int
    Next *ListNode
}


func PrintNode(list *ListNode) {
    print("[")
    for cur := list; cur != nil; cur = cur.Next {
        print(cur.Val, ",")
    }
    println("]")
}

func deleteDuplicates(head *ListNode) *ListNode {
    if nil == head || nil == head.Next {
        return head
    }

    cur := head.Val
    pos := head.Next
    node:= head

    for ;; {
        // 找到从 pos开始 第一个和 cur 不一样
        v := pos
        for ; nil != v && v.Val == cur; v = v.Next {}
        node.Next = v
        if nil == v {
            break
        }
        pos  = v
        node = node.Next
        cur  = v.Val
    }


    return head
}
