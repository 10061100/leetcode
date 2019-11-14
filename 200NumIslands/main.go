package main

import (
    "fmt"
)

func main() {
    fmt.Println(numIslands(convert([]string{"11110","11010", "11000", "00000"})))
}

func convert(ss []string) [][]byte{
    res := make([][]byte, len(ss))

    for i := 0; i < len(ss); i++ {
        res[i] = []byte(ss[i])
    }

    return res
}


var t = [][]int{{1,0},{-1,0},{0,1},{0,-1}}

func numIslands(grid [][]byte) int {
    hight := len(grid)
    if hight == 0 {
        return 0
    }

    width := len(grid[0])
    if width == 0 {
        return 0
    }

    visited := make([][]bool, hight)
    for i := 0; i < hight; i++ {
        visited[i] = make([]bool, width)
    }

    count := 0
    hq := NewQueue()
    wq := NewQueue()
    for i := 0; i < hight; i++ {
        for j := 0; j < width; j++ {
            if visited[i][j] == true {
                continue
            }

            if grid[i][j] == '0' {
                continue
            }

            // 否则, 从当前的开始去遍历
            hq.Push(i)
            wq.Push(j)
            count++
            visited[i][j] = true
            for ; !hq.IsEmpty(); {
                curh := hq.Front()
                curw := wq.Front()
                for _, di := range t {
                    newh := curh + di[0]
                    neww := curw + di[1]

                    if newh <0 || newh >= hight {
                        continue
                    }

                    if neww < 0 || neww >= width {
                        continue
                    }

                    if visited[newh][neww] == true {
                        continue
                    }

                    if grid[newh][neww] == '0' {
                        continue
                    }

                    visited[newh][neww] = true
                    hq.Push(newh)
                    wq.Push(neww)
                }
            }
        }
    }

    return count
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