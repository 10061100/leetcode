package main

func main() {

}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
    max := -111111

    maxPathSumSubOrder(root, &max)

    return max
}

func maxPathSumSubOrder(root *TreeNode, max *int) (rootSum int) {
    if root == nil {
        return 0
    }

    leftroot := maxPathSumSubOrder(root.Left, max)
    rightroot := maxPathSumSubOrder(root.Right, max)

    newroot:= Max(Max(root.Val, root.Val + leftroot), root.Val + rightroot)
    newval := Max(newroot, root.Val + leftroot + rightroot)

    if  newval > *max {
        *max = newval
    }

    return newroot
}



func Max(a, b int) int {
    if a > b {
        return a
    }

    return b
}