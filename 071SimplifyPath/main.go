package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println(simplifyPath("/aa/////") == "/aa")
    fmt.Println(simplifyPath("/home/") == "/home")
    fmt.Println(simplifyPath("/../") == "/")
    fmt.Println(simplifyPath("/home//foo") == "/home/foo")
    fmt.Println(simplifyPath("/a/./b/../../c/") == "/c")
}

func simplifyPath(path string) string {

    s := NewStack(100)
    for _, v := range strings.Split(path, "/") {
        v = strings.Trim(v, " ")

        if v == "" {
            continue
        }

        if v == "." {
            continue
        }

        if v == ".." {
            if !s.IsEmpty(){
                s.Pop()
            }
            continue
        }

        s.Push(v)
    }

    r := ""
    for i := 0; i < s.Len(); i++ {
        r = r + "/" + s.stack[i]
    }

    if r == "" {
        r = "/"
    }

    return r
}



type Stack struct {
    stack []string
    top int
    size int
}

func NewStack(n int) *Stack {
    s := make([]string, n)

    return &Stack{
        stack:s,
        top:0,
        size: n,
    }
}


func ( s *Stack) Top() string {
    if s.top == 0 {
        return ""
    }
    return s.stack[s.top-1]
}

func (s * Stack ) Pop() string {
    if s.top == 0 {
        return ""
    }

    s.top--
    r := s.stack[s.top]

    return r
}


func (s *Stack) Push(x string) *Stack {
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
