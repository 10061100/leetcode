package main

import (
    "fmt"
)

func main() {
    fmt.Println(removeDuplicates("aaabbba"))
}


func removeDuplicates(S string) string {
    bs := []byte(S)
    ls := len(bs)

    if ls <= 1 {
        return S
    }

    stack := NewStack(ls+1)
    // res := make([]byte, 0)
    for i := 0; i < ls; i++ {
        char := bs[i]
        if stack.IsEmpty() {
            stack.Push(char)
            continue
        }

        // fmt.Println(stack.Top(), ",", char, ",", stack.top)
        if stack.Top() == char {
            stack.Pop()
        } else {
            stack.Push(char)
        }
    }

    // fmt.Println(stack)
    return string(stack.stack[0:stack.top])
}


type Stack struct {
    stack []byte
    top int
    size int
}

func NewStack(n int) *Stack {
    s := make([]byte, n)

    return &Stack{
        stack:s,
        top:0,
        size: n,
    }
}


func ( s *Stack) Top() byte {
    if s.top == 0 {
        return 0
    }
    return s.stack[s.top-1]
}

func (s * Stack ) Pop() byte {
    if s.top == 0 {
        return 0
    }

    r := s.stack[s.top]
    s.top--

    return r
}


func (s *Stack) Push(x byte) *Stack {
    if s.top >= s.size {
        s.stack = append(s.stack, x)
        s.size++
    }


    s.stack[s.top] = x
    s.top++

    return s
}


func (s * Stack) IsEmpty() bool {
    return s.top == 0
}
