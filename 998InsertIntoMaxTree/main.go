package main


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


func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
    tnode := NewTreeNode(val)

    var pre *TreeNode = nil

    for node := root; ; {
        if node == nil {
            if pre != nil {
                pre.Right = tnode
            }
            break
        }

        if node.Val < val {
            if pre != nil {
                pre.Right = tnode
            }
            tnode.Left = node
            break
        }

        pre = node
        node = node.Right
    }

    if pre == nil {
        root = tnode
    }

    return root
}
