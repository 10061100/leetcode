package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func sufficientSubset(root *TreeNode, limit int) *TreeNode {

    t := dfs(root, limit)

    if !t {
        return nil
    }

    return root
}


func dfs(root *TreeNode, limit int) bool {
    if root.Left == nil && root.Right == nil {
        return root.Val >= limit
    }

    t := limit - root.Val
    leftok := false
    if root.Left != nil {
        leftok = dfs(root.Left, t)

    }
    rightok := false
    if root.Right != nil {
        rightok = dfs(root.Right, t)
    }

    if !leftok {
        root.Left = nil
    }

    if !rightok {
        root.Right = nil
    }

    return leftok || rightok
}
