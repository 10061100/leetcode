package main

import (
    "fmt"
)

func main()  {
    fmt.Println(longestWPI([]int{9,9,6,0,6,6,9}))
    fmt.Println(longestWPI([]int{6,9,9,6,9}))
}

func longestWPI(hours []int) int {
    dp := make([]int, len(hours) + 1)
    for i := 0; i < len(hours); i++ {
        dp[i+1] = -1
        if hours[i] > 8 {
            dp[i+1] = 1
        }
    }

    for i := 1; i < len(dp); i++ {
        dp[i] = dp[i-1] + dp[i]
    }

    // fmt.Println(dp)

    stack := NewStack(len(hours))
    for i := 0; i < len(dp); i++ {
        if stack.IsEmpty() {
            stack.Push(i)
            continue
        }

        if dp[i] < dp[stack.Top()] {
            stack.Push(i)
            continue
        }
    }

    // fmt.Println(stack.stack)

    ans := 0
    for i := len(dp) -1; i > 0; i-- {

        for ; !stack.IsEmpty() && dp[i] - dp[stack.Top()] > 0; {
            ans = max(ans, i - stack.Top())
            stack.Pop()
        }
    }

    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
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