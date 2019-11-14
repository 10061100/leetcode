package main

import (
    "fmt"
)

func main() {
    fmt.Println(shortestPathBinaryMatrix([][]int{{0,1},{1,0}}) == 2)
    fmt.Println(shortestPathBinaryMatrix([][]int{{0,0,0},{1,1,0},{1,1,0}}) == 4)
    fmt.Println(shortestPathBinaryMatrix([][]int{{0,0,0},{1,1,0},{1,1,0}}) == 4)
}

func shortestPathBinaryMatrix(grid [][]int) int {

    l := len(grid)

    if grid[0][0] == 1 || grid[l-1][l-1] == 1 {
        return -1
    }

    if l == 1 {
        return 1
    }

    xq := NewQueue()
    yq := NewQueue()

    t := make([][]int, l)
    for i := 0; i < l; i++ {
        t[i] = make([]int, l)
    }

    t[l-1][l-1] = 1
    xq.Push(l-1)
    yq.Push(l-1)

    direc := [][]int{{1,0},{1,1},{0,1},{-1,1},{-1,0},{-1,-1},{0,-1},{1,-1}}

    for ; !xq.IsEmpty(); {

        x := xq.Front()
        y := yq.Front()

        for _, v := range direc {
            newx := x + v[0]
            newy := y + v[1]

            if newx < 0 || newy < 0 {
                continue
            }

            if newx >= l || newy >= l {
                continue
            }

            // 过滤掉值为1的, 不能过去
            if grid[newx][newy] == 1 {
                continue
            }

            // 过滤掉已经访问的点
            if t[newx][newy] != 0 {
                continue
            }

            t[newx][newy] = 1 + t[x][y]
            xq.Push(newx)
            yq.Push(newy)
        }

    }


    if t[0][0] == 0 {
        return -1
    }

    return t[0][0]
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

    return q.head.Val
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