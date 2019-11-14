package main


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
    res := make([][]int, 0)

    levelOrderSub(root, 0, &res)

    return res
}

func levelOrderSub(root *TreeNode, level int, res *[][]int) {

    if root == nil {
        return
    }

    l := len(*res)

    if level > l-1 {
        *res = append(*res, make([]int, 0))
    }

    (*res)[level] = append((*res)[level], root.Val)

    levelOrderSub(root.Left, level+1, res)
    levelOrderSub(root.Right, level+1, res)
}