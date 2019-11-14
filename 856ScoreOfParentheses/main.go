package main

import (
    "fmt"
)

func main() {
    fmt.Println(scoreOfParentheses("()") == 1)
    fmt.Println(scoreOfParentheses("((()))") == 4)
    fmt.Println(scoreOfParentheses("(()())") == 4)
    fmt.Println(scoreOfParentheses("()()") == 2)
    fmt.Println(scoreOfParentheses("()(())(()())") == 7)
}

func scoreOfParentheses(S string) int {
    l := len(S)

    if l == 2 {
        return 1
    }

    stack := 0
    dp := make([]int, l/2+1)

    for _, c := range S {
        if c == '(' {
            // t := int(c)
            stack++
            continue
        }

        m := 1
        if dp[stack] != 0 {
            m = 2 * dp[stack]
            dp[stack] = 0
        }

        dp[stack-1] += m
        stack--
    }

    return dp[0]
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

    // s.stack[s.top] = 0
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