package main


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func flatten(root *TreeNode)  {
    if root == nil {
        return
    }

    last := flattenrc(root, nil)
    last.Right = nil
}


func flattenrc(root *TreeNode, pre *TreeNode) (last *TreeNode){

    left  := root.Left
    right := root.Right
    last   = root
    root.Left = nil
    root.Right = nil

    if pre != nil {
        pre.Right = root
    }

    if left != nil {
        last = flattenrc(left, root)
    }

    if right != nil {
        last = flattenrc(right, last)
    }

    return last
}
