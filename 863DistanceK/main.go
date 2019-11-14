package main

import (
    "fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func NewTreeNode(v int) *TreeNode {
    return &TreeNode{
        Val: v,
        Left: nil,
        Right: nil,
    }
}

func main() {

    root := NewTreeNode(1)

    root.Left = NewTreeNode(2)
    root.Right = NewTreeNode(3)
    root.Left.Right = NewTreeNode(4)
    root.Left.Left = NewTreeNode(5)
    root.Left.Right.Left = NewTreeNode(6)
    root.Left.Right.Right = NewTreeNode(7)

    fmt.Println(distanceK(root, root.Left.Right, 1))

}


func distanceK(root *TreeNode, target *TreeNode, K int) []int {
    // 分2个部分, 第一个部分是 当前target的子树上, 距离为k的值
    // 第二部分是和target拥有共同祖先其他分支上, 距离target的值
    res := make([]int, 0)
    // 第一部分
    dfs(target, 0, K, &res)

    // fmt.Println(res)
    // 第二部分
    // 先找到从root到target之间的所有共同祖先
    m := make(map[int]*TreeNode)
    stack := NewStack(10)
    rightVisited := make(map[int]bool)

    t := root
    isFirst := true

    for ; isFirst || !stack.IsEmpty(); {
        isFirst = false

        find := false
        for ; t != nil; {
            if t == target {
                find = true
                break
            }

            m[t.Val] = t
            stack.Push(t.Val)
            t = t.Left
        }

        // 如果已经找到target了, 就直接跳出
        if find {
            break
        }

        // 否则, 当前左子树已经完成, 从栈里面找一个, 右子树没有访问的
        for ; !stack.IsEmpty(); {
            tt := stack.Top()
            node := m[tt]
            // 如果右边已经访问了, 说明可以直接删除了
            if rightVisited[tt] != true && node.Right != nil {
                break
            }

            stack.Pop()
        }

        if stack.IsEmpty() {
            break
        }

        t = m[stack.Pop()]
        rightVisited[t.Val] = true
        t = t.Right
    }

    // 这里就找到了所有访问到了祖先

    pre := target
    deepth := 0

    for ; !stack.IsEmpty() && K - deepth >= 1 ; {
        t = m[stack.Pop()]
        deepth++

        if deepth == K {
            res = append(res, t.Val)
            break
        }

        if t.Left == pre {
            dfs(t.Right, 1, K-deepth, &res)
        } else {
            dfs(t.Left, 1, K-deepth, &res)
        }
    }


    return res
}


func dfs(node *TreeNode, deepth int, K int, res *[]int) {

    if node == nil {
        return
    }

    if deepth == K {
        *res = append(*res, node.Val)
        return
    }

    dfs(node.Left,  deepth+1, K, res)
    dfs(node.Right, deepth+1, K, res)
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
