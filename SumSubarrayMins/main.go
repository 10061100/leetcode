package main

func main()  {
    println(sumSubarrayMins([]int{3,1,2,4}))
}


type Pair struct {
    Val int
    Index int
}


func NewPair(v , i int) *Pair {
    p := new(Pair)
    p.Val = v
    p.Index = i

    return p
}


type Stack struct {
    top int
    stack []*Pair
}

func NewStack(len int) *Stack {
    return &Stack{
        top: -1,
        stack: make([]*Pair, len),
    }
}

func (s *Stack) IsEmpty() bool {
    return s.top < 0
}

func (s *Stack) GetTopPair() *Pair {
    if s.IsEmpty() {
        return nil
    }

    return s.stack[s.top]
}

func (s *Stack) Reset() {
    s.top = -1
}

func (s *Stack) Pop() *Pair {
    if s.top < 0 {
        return nil
    }
    t := s.stack[s.top]
    s.top--
    return t
}


func (s *Stack) Push(p *Pair) {
    s.top++
    s.stack[s.top] = p
}

func sumSubarrayMins(A []int) int {
    la := len(A)
    if la == 1 {
        return A[0]
    }

    // 本质上在找左边&右边第一个比自己小的数
    left := findLeftMinIndex(A)
    right:= findRightMinIndex(A)

    // fmt.Println(left)
    // fmt.Println(right)

    r := 0
    for i, v  := range A {
        l := left[i]
        rr := right[i]

        ll :=  (i-l+1) * (rr - i + 1)
        r = (r + ( v * ll ) % 1000000007) % 1000000007
    }

    return r
}


func findLeftMinIndex(A []int) []int {
    r := make([]int, len(A))
    stack := NewStack(len(A))
    for i, v := range A {
        for ; ; {
            if stack.IsEmpty() {
                r[i] = 0
                stack.Push(NewPair(v, i))
                break
            }

            if top := stack.GetTopPair(); top.Val <= v {
                r[i] = top.Index + 1
                stack.Push(NewPair(v, i))
                break
            }

            stack.Pop()
        }
    }

    return r
}


func findRightMinIndex(A []int) []int {
    r := make([]int, len(A))
    stack := NewStack(len(A))
    for i := len(A)-1; i >= 0; i-- {
        v := A[i]
        for ; ; {
            if stack.IsEmpty() {
                r[i] = len(A)-1
                stack.Push(NewPair(v, i))
                break
            }

            if top := stack.GetTopPair(); top.Val < v {
                r[i] = top.Index - 1
                stack.Push(NewPair(v, i))
                break
            }

            stack.Pop()
        }
    }

    return r
}