package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}



func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return NewTreeNode(val)
    }

    if val > root.Val {
        root.Right = insertIntoBST(root.Right, val)
    } else {
        root.Left = insertIntoBST(root.Left, val)
    }

    return root
}


func NewTreeNode(v int) *TreeNode {
    return &TreeNode{
        Val: v,
        Left: nil,
        Right: nil,
    }
}
