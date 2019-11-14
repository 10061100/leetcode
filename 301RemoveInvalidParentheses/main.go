package main



func removeInvalidParentheses(s string) []string {
    bs := []byte(s)
    res := make([]string, 0)

    if len(bs) == 0 {
        return res
    }

    vq := NewStack(len(bs))
    // iq := NewStack(len(bs))
    // preRight := -1

    for i := 0; i < len(bs); i++ {
        n := bs[i]

        if n != '(' && n != ')' {
            continue
        }

        if n == '(' {
            vq.Push(i)
        } else {
            if vq.IsEmpty() {
                // 如果栈底是空的, 说明这个')'是多余的, 这个括号是需要去除的
            } else {
                vq.Pop()
            }

            // preRight = i
        }
    }

    return res
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

    r := s.stack[s.top]
    s.top--

    return r
}


func (s *Stack) Push(x int) *Stack {
    if s.top >= s.size {
        s.stack = append(s.stack, x)
    }

    s.top++
    s.stack[s.top] = x

    return s
}


func (s * Stack) IsEmpty() bool {
    return s.top == 0
}