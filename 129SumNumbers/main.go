package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}




func sumNumbers(root *TreeNode) int {
    sum := 0
    pre := 0

    dfs(root, pre, &sum)

    return sum
}

func dfs(node *TreeNode, pre int, sum *int)  {


    if node == nil {
        return
    }

    t := (pre * 10) + node.Val

    if node.Left == nil && node.Right == nil {
        *sum = *sum + t
        return
    }

    dfs(node.Left, t, sum)
    dfs(node.Right, t, sum)

}
