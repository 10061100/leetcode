package main

import (
    "fmt"
)

func main() {
    fmt.Println(maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7}, 3))
    fmt.Println(maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7}, 1))
    fmt.Println(maxSlidingWindow([]int{1,3,1,2,0,5}, 3))
}


func maxSlidingWindow(nums []int, k int) []int {
    res := make([]int, 0)
    l := len(nums)

    if l == 0 {
        return res
    }

    valq := NewQueue()
    idxq := NewQueue()

    valq.Push(nums[0])
    idxq.Push(0)
    if k == 1 {
        res = append(res, valq.First())
    }

    // 初始化
    for i := 1; i < l ; i++ {
        n := nums[i]

        // 先判断是否需要进行首部出队列操作
        if i >= k {
            // 需要先进行pop操作
            idx := idxq.First()
            if idx == i-k {
                // 如果当前值是对首的, 则需要先出库
                idxq.Front()
                valq.Front()
            }
        }

        // 之后进行对首小的,就把他们都出队列, 然后自己进去
        for ; !valq.IsEmpty() && valq.Last() <= n; {
            valq.Pop()
            idxq.Pop()
        }

        // 把自己放进队列
        valq.Push(n)
        idxq.Push(i)

        if i >= k -1 {
            res = append(res, valq.First())
        }
    }

    return res
}


type QueueNode struct {
    Val int
    Next *QueueNode
    Pre  *QueueNode
}

func NewQueueNode(v int) *QueueNode{
    return &QueueNode{
        Val: v,
        Next: nil,
        Pre : nil,
    }
}

type Queue struct {
    head *QueueNode
    tail *QueueNode
}


func NewQueue() *Queue{
    return &Queue{
        head: nil,
        tail: nil,
    }
}


func (q *Queue) Push(val int) *Queue {
    node := NewQueueNode(val)

    if q.tail == nil {
        q.head = node
        q.tail = q.head
    } else {
        q.tail.Next = node
        node.Pre = q.tail
        q.tail = node
    }

    return q
}

func (q *Queue) IsEmpty() bool {
    return q.head == nil
}

func (q *Queue) First() int {
    if q.IsEmpty() {
        return 0
    }

    return q.head.Val
}

func (q *Queue) Front() int {
    if q.IsEmpty() {
        return 0
    }

    t := q.head
    q.head = q.head.Next
    if q.head != nil {
        q.head.Pre = nil
    }
    if q.head == nil {
        q.tail = nil
    }

    return t.Val
}

func( q *Queue) Last() int {
    if q.IsEmpty() {
        return 0
    }

    return q.tail.Val
}

func ( q* Queue) Pop() int {
    if q.IsEmpty() {
        return 0
    }

    node := q.tail

    q.tail = node.Pre
    if q.tail != nil {
        q.tail.Next = nil
    }
    if q.tail == nil {
        q.head = nil
    }

    return node.Val
}



