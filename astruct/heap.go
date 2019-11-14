package astruct


type HeapInterface interface {
    Compare(t HeapInterface) bool
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

    for j := len(h.h)-1; j > 0; j-- {
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

func (h *Heap) AddPop(t HeapInterface) HeapInterface {
    if len(h.h) == 0 {
        return nil
    }

    res := h.h[0]

    if t == nil {
        h.h[0] = h.h[len(h.h)-1]
        h.h = h.h[0:len(h.h) - 1]
    } else {
        h.h[0] = t
    }


    h.Adjust(0)

    return res
}