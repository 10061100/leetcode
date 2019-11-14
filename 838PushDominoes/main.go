package main

import (
    "fmt"
)

func main() {
    fmt.Println(pushDominoes(".") == ".")
    fmt.Println(pushDominoes("L") == "L")
    fmt.Println(pushDominoes("R") == "R")
    fmt.Println(pushDominoes("LR") == "LR")
    fmt.Println(pushDominoes("R.L") == "R.L")
    fmt.Println(pushDominoes("R...L") == "RR.LL")
    fmt.Println(pushDominoes(".L.R...LR..L..") == "LL.RR.LLRRLL..")
    fmt.Println(pushDominoes("RRRRRRL...")=="RRRRRRL...")

}


func pushDominoes(dominoes string) string {
    if dominoes == "" {
        return ""
    }

    stack := NewStack(len(dominoes))
    left2right := make([]int, len(dominoes))
    res := make([]byte, len(dominoes))

    // 从左向右遍历, 寻找每个位置左边第一个向右的坐标, 没有就给-1
    for k, c := range dominoes {
        if c == 'R' {
            stack.Push(k)
        } else if c == 'L' {
            // 向左的话, 无所谓是什么, 因为下一轮肯定是向左
            stack.Reset()
        }

        if stack.IsEmpty() {
            left2right[k] = -1
        } else {
            left2right[k] = stack.Top()
        }
    }

    stack.Reset()

    // 从右向左遍历, 找到每个位置右边第一个向左倾斜的动作
    for i := len(dominoes)-1; i >= 0; i-- {
        c := dominoes[i]

        if c == 'L' {
            stack.Push(i)
        } else if c == 'R' && !stack.IsEmpty() {
            stack.Reset()
        }

        n := -1
        if !stack.IsEmpty() {
            n = stack.Top()
        }

        if n == -1 && left2right[i] == -1 {
            res[i] = '.'
        } else if n == -1 {
            res[i] = 'R'
        } else if left2right[i] == -1 {
            res[i] = 'L'
        } else {
            left := abs(left2right[i] - i)
            right:= abs(n - i)

            // fmt.Println(left2right[i], ",", i, ",", n)

            if left > right {
                res[i] = 'L'
            } else if left < right {
                res[i] = 'R'
            } else {
                res[i] = '.'
            }
        }
    }


    return string(res)
}

func abs(a int) int {
    if a >= 0 {
        return a
    }

    return -a
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
    if !s.IsEmpty() {
        s.top = 0
    }
}