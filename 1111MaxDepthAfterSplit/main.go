package main

import (
    "fmt"
)

func main()  {
    fmt.Println(maxDepthAfterSplit("(()())"))
    fmt.Println(maxDepthAfterSplit("()(())()"))
}


func maxDepthAfterSplit(seq string) []int {
    l := len(seq)

    res := make([]int, l)

    if l == 0 {
        return res
    }

    stack := NewStack(l/2)

    for i, v := range []byte(seq) {
        if v == '(' {
            stack.Push(i)
            // res[i] = (stack.top + 1)%2
        } else {
            t := stack.top
            top := stack.Pop()
            res[i] , res[top] = (t+1)%2, (t+1)%2
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

func (this *Stack) ToArray() []int {
    return this.stack[0:this.top]
}
