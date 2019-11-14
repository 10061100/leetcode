package main


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}





func addOneRow(root *TreeNode, v int, d int) *TreeNode {
    if d == 1 {
        p := NewTree(v)
        p.Left = root
        root = p
    } else {
        dfs(root, 1, v, d)
    }

    return root
}


func dfs(node *TreeNode, deepth int, v int, d int) {
    if deepth == d - 1 {
        left := node.Left
        right:= node.Right

        if left != nil {
            p := NewTree(v)
            p.Left = left
            left = p
            node.Left = left
        }

        if right != nil {
            p := NewTree(v)
            p.Right = right
            right = p
            node.Right = right
        }

        return
    }

    if node.Left != nil {
        dfs(node.Left, deepth+1, v, d)
    }

    if node.Right != nil {
        dfs(node.Right, deepth+1, v, d)
    }
}


func NewTree(v int) *TreeNode {
    return &TreeNode{
        Val: v,
        Left: nil,
        Right: nil,
    }
}