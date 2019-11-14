package main

import (
    "fmt"
    "sort"
)


func main() {
    fmt.Println(advantageCount([]int{2,7,11,15}, []int{1,10,4,11}))
    fmt.Println(advantageCount([]int{12,24,8,32}, []int{13,25,32,11}))
}

func advantageCount(A []int, B []int) []int {
    sortedA := make([]int, len(A))
    sortedB := make([]int, len(B))

    copy(sortedA, A)
    copy(sortedB, B)

    sort.Sort(sort.IntSlice(sortedA))
    sort.Sort(sort.IntSlice(sortedB))

    // 对于每一个B中的元素, 查看B的位置排序
    valid := make(map[int]*Stack)
    invalid := NewStack(len(A))

    i := 0
    for _, v := range sortedA {
        // 说明当前的数可以当做是b这个数的备用
        if v > sortedB[i] {
            if _, ok := valid[sortedB[i]]; !ok {
                valid[sortedB[i]] = NewStack(len(A))
            }
            valid[sortedB[i]].Push(v)
            i++
        } else {
            invalid.Push(v)
        }
    }

    res := make([]int, 0)
    for _ , v := range B {
        if t, ok := valid[v]; !ok || t.IsEmpty() {
            res = append(res, invalid.Pop())
        } else {
            res = append(res, t.Pop())
        }
    }

    return res
}



type Stack struct {
    stack []int
    top int
    size int
}

func NewStack(n int) *Stack {
    s := make([]int, n)

    return &Stack{
        stack:s,
        top:0,
        size: n,
    }
}


func ( s *Stack) Top() int {
    if s.top == 0 {
        return 0
    }
    return s.stack[s.top-1]
}

func (s * Stack ) Pop() int {
    if s.top == 0 {
        return 0
    }

    s.top--
    r := s.stack[s.top]

    return r
}


func (s *Stack) Push(x int) *Stack {
    if s.top >= s.size {
        s.stack = append(s.stack, x)
    }

    s.stack[s.top] = x
    s.top++

    return s
}


func (s * Stack) IsEmpty() bool {
    return s.top == 0
}

func (s * Stack) Len() int {
    return s.top
}

func (s *Stack) Reset() {
    for ; !s.IsEmpty(); {
        s.Pop()
    }
}