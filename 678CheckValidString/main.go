package main

import (
    "fmt"
)

func main() {
    fmt.Println(checkValidString("**((*"))
    fmt.Println(checkValidString("*)(*"))
    fmt.Println(checkValidString("(())((())()()(*)(*()(())())())()()((()())((()))(*"))
    fmt.Println(checkValidString("(*))"))
}

func checkValidString(s string) bool {
    stack := NewStack(len(s))
    star  := NewStack(len(s))
    for k, v := range s {
        if v == '(' {
            stack.Push(k)
        } else if v == '*' {
            star.Push(k)
        } else {
            if stack.IsEmpty() && star.IsEmpty() {
                return false
            } else if !stack.IsEmpty() {
                stack.Pop()
            } else {
                star.Pop()
            }
        }
    }

    count := 0
    for ; !stack.IsEmpty() ; {
        top := stack.Top()
        for ; !star.IsEmpty() && star.Top() > top; {
            count++
            star.Pop()
        }

        if count <= 0 {
            return false
        }

        stack.Pop()
        count--
    }

    return true
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