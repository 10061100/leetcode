package main


type Queue struct {
    queue []int
    size  int
}

func NewQueue() *Queue {
    return &Queue{
        queue: make([]int, 0),
        size : 0,
    }
}

func (q *Queue) IsEmpty() bool {
    return q.size == 0
}

func (q *Queue) Add(a int) {
    q.queue = append(q.queue, a)
    q.size++
}

func (q *Queue) Out() int {
    if q.IsEmpty() {
        return -1
    }

    t := q.queue[0]
    q.queue = q.queue[1:]
    q.size--

    return t
}

func (q *Queue) First() int {
    if q.IsEmpty() {
        return -1
    }

    return q.queue[0]
}

func (q * Queue) Last() int {
    if q.IsEmpty() {
        return -1
    }

    return q.queue[q.size-1]
}


func shortestSubarray(A []int, K int) int {
    s := len(A)
    if s == 0 {
        return -1
    }

    tmp := make([]int, s)
    tmp[0] = A[0]
    for i := 1 ; i < s; i++ {
        tmp[i] = tmp[i-1] + A[i]
    }

    q := NewQueue()

    q.Add(tmp[0])
    for i := 1; i < s; i++ {
        if q.IsEmpty() {
            q.Add(tmp[i])
            continue
        }

        if tmp[i] < q.First() {

        }

    }

}


func main() {



}


