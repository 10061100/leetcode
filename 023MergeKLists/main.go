package main


type ListNode struct {
    Val int
    Next *ListNode
}


func mergeKLists(lists []*ListNode) *ListNode {
    l := len(lists)
    if l == 0 {
        return nil
    }

    if l == 1 {
        return lists[0]
    }


}
