package main

import (
    "fmt"
)

func main() {

    fmt.Println(kthSmallest([][]int{{1,2},{2,3}}, 1) == 1)
    fmt.Println(kthSmallest([][]int{{1,2},{2,3}}, 2) == 2)
    fmt.Println(kthSmallest([][]int{{1,2},{2,3}}, 3) == 2)
    fmt.Println(kthSmallest([][]int{{1,2},{2,3}}, 4) == 3)
    fmt.Println(kthSmallest([][]int{{1,5,9},{10,11,13},{12,13,15}}, 4) == 10)
}


func kthSmallest(matrix [][]int, k int) int {
    n := len(matrix)
    h := NewHeap()
    i := 0
    for ; i < n && i < k ; i++ {
        h.Add([]HeapInterface{NewT(matrix[i][0], i*n)})
    }

    h.Init()
    for ; k > 1 ; k-- {
        var change *T = nil

            // 说明需要再进一个
         min := h.Top().(*T)
         if t := min.Index % n; t < n-1 {
            change = NewT(matrix[min.Index/n][(t+1)], min.Index+1)
         }

        // fmt.Println("xxxxxx", change.GetVal())
        h.AddPop(change, change == nil)

    }

    return h.Top().GetVal()
}


type T struct {
    t int
    Index int
}

func NewT(v, index int) *T{
    return &T{
        t : v,
        Index: index,
    }
}

func( t *T) GetVal() int {
    return t.t
}

func(t *T) Compare(m HeapInterface) bool{
    if t.GetVal() > m.GetVal() {
        return true
    }

    return false
}



type HeapInterface interface {
    Compare(t HeapInterface) bool
    GetVal() int
}

type Heap struct {
    h []HeapInterface
    size int
}


func NewHeap() *Heap {
    return &Heap{
        h: make([]HeapInterface, 0),
    }
}

func (h *Heap) Add(t []HeapInterface) *Heap {
    h.h = append(h.h, t...)

    return h
}


func (h *Heap) Init() *Heap {

    for j := len(h.h)-1; j >= 0; j-- {
        h.Adjust(j)
    }

    return h
}

// 调整以i为根的子堆
func (h *Heap) Adjust(i int) {

    if (2*i + 1) >= len(h.h) {
        return
    }

    min, minindex := h.h[2*i + 1], 2*i+1
    // 如果左子树比右子树大, 说明右子树是小的
    if (2*i + 2) < len(h.h) && min.Compare(h.h[2*i+2]){
        min = h.h[2*i+2]
        minindex = 2*i + 2
    }

    // 如果较小的子树, 比自己还小, 则说明需要调节当前节点
    if !min.Compare(h.h[i]) {
        h.h[i], h.h[minindex] = min, h.h[i]
        h.Adjust(minindex)
    }
}

func (h *Heap) Top() HeapInterface {
    if len(h.h) == 0 {
        return nil
    }

    return h.h[0]
}

func (h *Heap) AddPop(t HeapInterface, isNil bool) HeapInterface {
    if len(h.h) == 0 {
        return nil
    }

    res := h.h[0]

    // fmt.Println(h.h)
    // fmt.Println(t == nil)
    if t == nil || isNil {
        h.h[0] = h.h[len(h.h)-1]
        h.h = h.h[0:len(h.h) - 1]
    } else {
        h.h[0] = t
    }

    //fmt.Println(h.h)
    h.Adjust(0)

    return res
}