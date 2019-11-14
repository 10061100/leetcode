package main

import (
    "fmt"
    "strconv"
)

func main() {
    fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}) == 9)
    fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}) == 6)
    fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}) ==22)
}

func evalRPN(tokens []string) int {
    stack := NewStack(len(tokens))

    for i := 0; i < len(tokens); i++ {
        switch tokens[i] {
        case "+":
            stack.Push(stack.Pop() + stack.Pop())
        case "-":
            t1 := stack.Pop()
            t2 := stack.Pop()
            stack.Push(t2-t1)
        case "*":
            stack.Push(stack.Pop() * stack.Pop())
        case "/":
            t1 := stack.Pop()
            t2 := stack.Pop()
            stack.Push(t2/t1)
        default:
            v, _ := strconv.Atoi(tokens[i])
            stack.Push(v)
        }
    }

    return stack.Pop()
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