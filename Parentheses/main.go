
// Parentheses
package main

func main() {
    println(isValid("{}{}{}{}{}{}{}{}{}{}"))
}


func isValid(s string) bool {
    stack := NewStack(len(s)/2 + 1)
    for _, c := range s {
        sb := getSymble(c)
        if sb > 0 {
            // 左括号, 可以入栈
            if !stack.Push(sb) {
                return false
            }
        } else {
            if stack.Pop() + sb != 0 {
                return false
            }
        }
    }


    return stack.IsEmpty()
}


type Stack struct {
    stack []int
    top int
    size int
}


func NewStack(n int) *Stack {
    s := make([]int, n)

    for i := 0; i < n ; i++ {
        s = append(s, 0)
    }

    return &Stack{
        stack:s,
        top:0,
        size: n,
    }
}


func (s * Stack ) Pop() int {
    if s.top == 0 {
        return 0
    }

    r := s.stack[s.top]
    s.top--

    return r
}


func (s *Stack) Push(x int) bool {
    if s.top >= s.size {
        return false
    }

    s.top++
    s.stack[s.top] = x

    return true
}


func (s * Stack) IsEmpty() bool {
    return s.top == 0
}


func getSymble(c rune) int {
    switch c {
    case '(':
        return 1
    case ')':
        return -1
    case '[':
        return 2
    case ']':
        return -2
    case '{':
        return 3
    case '}':
        return -3
    default:
        return 0
    }
}
