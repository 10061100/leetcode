package astruct

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

func midTravel(root *TreeNode) {
    if root == nil {
        return
    }

    if root.Left != nil {
        midTravel(root.Left)
    }

    fmt.Println(root.Val)

    if root.Right != nil {
        midTravel(root.Right)
    }

}