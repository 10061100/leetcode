package main


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
    res := 0
    length := 1

    findSub(root, 1, &res, &length)

    return res
}

func findSub(root *TreeNode, cl int, val *int, length *int) {
    if root == nil {
        return
    }

    if cl > *length {
        *length = cl
        *val = root.Val
    }

    findSub(root.Left, cl+1, val, length)
    findSub(root.Right, cl+1, val, length)
}
