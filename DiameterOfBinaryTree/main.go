package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


func main() {
    println(diameterOfBinaryTree(nil))
}

func diameterOfBinaryTree(root *TreeNode) int {
    max := 0
    maxLength(root, &max)

    return max
}


func maxLength(root *TreeNode, max *int) int {

    if root == nil {
        return 0
    }

    left := maxLength(root.Left, max)
    right:= maxLength(root.Right, max)

    nl := left + right + 1

    if nl > *max {
        *max = nl
    }

    if left > right {
        return left + 1
    }

    return right + 1
}
