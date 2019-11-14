package main

import (
    "fmt"
)

func main() {
    s := Constructor()

    for _, v := range []int{1,79, 34, 21,34, 16} {
        fmt.Println(s.Next(v))
    }
}


type StockSpanner struct {
    price *Stack
    index *Stack
    curIndex int
}

const MAX = 100001


func Constructor() StockSpanner {
    return StockSpanner{
        price: NewStack(10001).Push(MAX),
        index: NewStack(10001).Push(0),
        curIndex:0,
    }
}


func (this *StockSpanner) Next(price int) int {
    this.curIndex++
    for ; !this.price.IsEmpty() && price >= this.price.Top(); {
        this.price.Pop()
        this.index.Pop()
    }

    res := this.curIndex - this.index.Top()

    if price < this.price.Top() {
        this.price.Push(price)
        this.index.Push(this.curIndex)
    }

    // fmt.Println(this.price.ToArray())
    // fmt.Println(this.index.ToArray())

    return res
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */

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

func( s *Stack) ToArray() []int {
    return s.stack[0:s.top]
}