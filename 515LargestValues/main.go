package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func largestValues(root *TreeNode) []int {
    res := make([]int, 0)

    if root == nil {
        return res
    }

    dfs(root, 0, &res)

    return res
}

func dfs(root *TreeNode, deep int, res *[]int) {

    if root == nil {
        return
    }

    if deep >= len(*res) {
        *res = append(*res, root.Val)
    } else {
        if t := (*res)[deep]; t < root.Val {
            (*res)[deep] = root.Val
        }
    }

    dfs(root.Left, deep+1, res)
    dfs(root.Right, deep+1, res)

}



