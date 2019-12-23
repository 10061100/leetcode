package main


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}




func bstToGst(root *TreeNode) *TreeNode {
    sum := 0

    subBstToGst(root, &sum)

    return root
}


func subBstToGst(root *TreeNode, sum *int) {
    if root == nil {
        return
    }

    subBstToGst(root.Right, sum)

    root.Val += *sum

    *sum = root.Val

    subBstToGst(root.Left, sum)
}

