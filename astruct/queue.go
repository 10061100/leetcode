package astruct


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