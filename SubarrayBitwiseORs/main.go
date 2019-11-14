package main


func main() {
    println(subarrayBitwiseORs([]int{1,2,4}))

    // isValidBST2()
}

func subarrayBitwiseORs(A []int) int {
    if len(A) <= 1 {
        return len(A)
    }

    // max := math.MaxInt32
    m := make(map[int]int)

    dp := make([]int, len(A))
    for i, v := range A {
        m[v] = 1
        dp[i] = v
    }

    l := len(A)
    for i := 1; i < l; i++ {
        for j := 0; j + i < l; j++ {
            dp[j] = dp[j] | A[j+i]
            m[dp[j]] = 1
        }
    }

    return len(m)
}




type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
    if root == nil {
        return true
    }

    if root.Left == nil && root.Right == nil {
        return true
    }

    if root.Left == nil || root.Right == nil {
        return false
    }

    if root.Left.Val >= root.Val || root.Right.Val <= root.Val {
        return false
    }

    if !isValidBST(root.Left) || !isValidBST(root.Right) {
        return false
    }

    max := maxNode(root.Left)
    if max.Val >= root.Val {
        return false
    }

    min := minNode(root.Right)
    if min.Val <= root.Val {
        return false
    }

    return true
}


func maxNode(root *TreeNode) *TreeNode{
    t := root
    for ; t.Right != nil ; t = t.Right {}

    return t
}

func minNode(root *TreeNode) *TreeNode{
    t := root
    for ; t.Left != nil ; t = t.Left {}

    return t
}


func isValidBSTWithLast(root *TreeNode, pre *TreeNode) (bool, *TreeNode) {

    if root == nil {
        return true, nil
    }

    if root.Left != nil {
        isbst := false
        isbst, pre = isValidBSTWithLast(root.Left, pre)
        if !isbst {
            return false, root
        }
    }

    if pre != nil && pre.Val >= root.Val {
        return false, nil
    }

    retlast := root
    if root.Right != nil {
        isbst := false
        isbst, retlast = isValidBSTWithLast(root.Right, root)
        if !isbst {
            return false, root
        }
    }

    return true, retlast

}



func isValidBST2(root *TreeNode) bool {
    s := NewStack(10)

    if root == nil {
        return false
    }

    s.Push(root)

    var pre *TreeNode = nil
    for ; !s.IsEmpty(); {
        top := s.Top()
        for ; top.Left != nil; top = top.Left {
            s.Push(top.Left)
        }

        top = s.Pop()
        if pre != nil && pre.Val > top.Val {
            return false
        }

        pre = top
        if top.Right != nil {
            s.Push(top.Right)
        }
    }

    return true

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
        s.size++
    }

    s.top++
    s.stack[s.top] = x

    return s
}


func (s * Stack) IsEmpty() bool {
    return s.top == 0
}