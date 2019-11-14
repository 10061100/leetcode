package main


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
    res := make([]int, 0)
    if root == nil {
        return res
    }

    st := NewStack(100)
    st.Push(root)


    for ; !st.IsEmpty(); {
        // 一定不是nil
        top := st.Pop()
        for ; top != nil; {
            st.Push(top)
            top = top.Left
        }

        // 这个时候top已经是nil了
        top = st.Pop()
        // 遇到了一个节点这个节点没有左子树了
        res = append(res, top.Val)

        if top.Right != nil {
            st.Push(top.Right)
        }

    }

    return res
}

type Stack struct {
    stack []*TreeNode
    top int
    size int
}

func NewStack(n int) *Stack {
    s := make([]*TreeNode, n)

    return &Stack{
        stack:s,
        top:0,
        size: n,
    }
}


func ( s *Stack) Top() *TreeNode {
    if s.top == 0 {
        return nil
    }
    return s.stack[s.top-1]
}

func (s * Stack ) Pop() *TreeNode {
    if s.top == 0 {
        return nil
    }

    r := s.stack[s.top]
    s.top--

    return r
}


func (s *Stack) Push(x *TreeNode) *Stack {
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
