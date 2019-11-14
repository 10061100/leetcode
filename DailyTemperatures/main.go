package main

import (
    "fmt"
)

func main() {
    fmt.Println(dailyTemperatures([]int{89,62,70,58,47,47,46,76,100,70}))
}

func dailyTemperatures(T []int) []int {
    res := make([]int, len(T))

    vs := NewStack(len(T))
    is := NewStack(len(T))
    res[len(T)-1] = 0

    vs.Push(T[len(T)-1])
    is.Push(len(T)-1)

    for i := len(T)-2; i >= 0; i-- {
        for ; !vs.IsEmpty() && vs.Top() <= T[i]; {
            vs.Pop()
            is.Pop()
        }

        if vs.IsEmpty() {
            res[i] = 0
        } else {
            res[i] = is.Top() - i
        }

        vs.Push(T[i])
        is.Push(i)
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

    for i := 0; i < n ; i++ {
        s[i] = 0
    }

    return &Stack{
        stack:s,
        top:0,
        size: n,
    }
}

func( s *Stack) Top() int {
    return s.stack[s.top]
}


func (s * Stack ) Pop() int {
    if s.top == 0 {
        return 0
    }

    r := s.stack[s.top]
    s.top--

    return r
}


func (s *Stack) Push(x int) bool {
    if s.top >= s.size {
        return false
    }

    s.top++
    s.stack[s.top] = x

    return true
}


func (s * Stack) IsEmpty() bool {
    return s.top == 0
}