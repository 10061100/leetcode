package main

import (
    "fmt"
    "sort"
)

func main() {
    fmt.Println(maxWidthRamp([]int{6,0,8,2,1,5}) == 4)
    fmt.Println(maxWidthRamp([]int{9,8,1,0,1,9,4,0,4,1}) == 7)
}

func maxWidthRamp(A []int) int {

    if len(A) <= 1 {
        return 0
    }

    stack := NewStack(len(A))
    index := NewStack(len(A))

    res := 0
    for k, v := range A {
        if stack.IsEmpty() {
            stack.Push(v)
            index.Push(k)
            continue
        }

        if stack.Top() > v {
            stack.Push(v)
            index.Push(k)
            continue
        }

        // fmt.Println(index.stack[0:stack.top])
        tk := sort.Search( stack.top, func(i int) bool {
            if stack.stack[i] > v {
                return false
            }

            return true
        })

        if tk == stack.top {
            continue
        }

        // fmt.Println(k, "====", stack.stack[k])

        res = max(res, k - index.stack[tk])
    }

    return res

}

func max(i, j int) int {
    if i > j {
        return i
    }

    return j
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