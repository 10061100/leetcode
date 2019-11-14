package main

import (
    "fmt"
)

func main() {
    fmt.Println(removeDuplicateLetters("cbabc"))
}

func removeDuplicateLetters(s string) string {
    c := [26]int{0}
    for i := len(s)-1; i >= 0; i-- {
        n := s[i] - 'a'
        c[n]++
    }

    stack := NewStack(len(s))

    for _, v := range []byte(s) {
        n := int(v - 'a')
        c[n]--
        if stack.IsEmpty() {
            stack.Push(n)
        } else {
            if stack.IsExist(n) {
                // 如果已经存在, 直接就过去了
                continue
            }
            // 如果不存在, 就把栈顶所有符合需求的都出栈
            // 符合出栈需求: 1. 后面还会出现; 2 比当前的字母要大
            for ; !stack.IsEmpty(); {
                topc := stack.Top()
                if c[topc] > 0 && topc > n {
                    stack.Pop()
                } else {
                    break
                }
            }

            stack.Push(n)
        }
    }

    res := make([]byte, stack.Len())
    l := stack.Len() - 1
    for i := 0 ; !stack.IsEmpty(); i++ {
        res[l-i] = byte(stack.Pop() + 'a')
    }

    return string(res)
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

func (s *Stack) IsExist(n int) bool {
    for i := 0; i < s.top; i++ {
        if n == s.stack[i] {
            return true
        }
    }

    return false
}

