// LongestUnivaluePath
package main

func main()  {
    println(longestUnivaluePath(nil))
}

type Queue struct {
    C []*TreeNode
    N int
}

func NewQueue() *Queue {
    return &Queue{
        C: make([]*TreeNode, 0),
        N: 0,
    }
}

func (q *Queue) Push(t *TreeNode){
    q.C = append(q.C, t)
    q.N++
}

func (q *Queue) IsEmpty() bool {
    return q.N == 0
}

func (q *Queue) Out() *TreeNode {
    if q.IsEmpty() {
        return nil
    }

    q.N--
    t := q.C[0]
    q.C = q.C[1:]
    return t
}

func longestUnivaluePath(root *TreeNode) int {
    if root == nil {
        return 0
    }
    q := NewQueue()
    q.Push(root)
    max := 0
    for ; !q.IsEmpty(); {
        t := q.Out()
        if t == nil {
            break
        }

        if l := pathLength(t, q); l > max {
            max = l
        }

    }

    return max - 1
}

func pathLength(root * TreeNode, q *Queue) int {
    left := 0
    right:= 0

    if root.Left != nil {
        if root.Left.Val == root.Val {
            left = rootLength(root.Left, q)
        } else {
            q.Push(root.Left)
        }
    }

    if root.Right != nil {
        if root.Right.Val == root.Val {
            right = rootLength(root.Right, q)
        } else {
            q.Push(root.Right)
        }
    }

    return left + right + 1
}


func rootLength(root *TreeNode, q *Queue) int {
    c := 1;
    left := 0
    right:= 0
    if root.Left != nil {
        if root.Left.Val == root.Val {
            left = rootLength(root.Left, q)
        } else {
            q.Push(root.Left)
        }
    }

    if root.Right != nil {
        if root.Right.Val != root.Val {
            q.Push(root.Right)
        } else {
            right = rootLength(root.Right, q)
        }
    }

    if root.Left == nil && root.Right == nil {
        return c
    }

    if left > right {
        return c + left
    }

    return c + right
}


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

