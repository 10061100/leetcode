package main

func find132pattern(nums []int) bool {
    n := len(nums)

    if n <= 2 {
        return false
    }

    // 寻找每个数左边比当前数大的最近的位置
    leftbig     := make([]int, n)
    // 寻找每个数左边比当前数小的最远的位置
    leftsmall   := make([]int, n)

    bstack := NewStack(n)
    sstack := NewStack(n)
    leftbig[0] = -1
    leftsmall[0] = -1
    bstack.Push(nums[0])
    sstack.Push(0)

    // 寻找左边比自己小的数的最远位置
    for i := 1; i < n; i++ {

    }

    return false
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