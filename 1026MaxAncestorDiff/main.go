package main


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
    a := 0
    b := 0

    return dfs(root, &a, &b)
}


func dfs(n *TreeNode, mmin, mmax *int ) int {
    if n == nil {
        return 0
    }

    *mmin = min(*mmin, n.Val)
    *mmax = max(*mmax, n.Val)

    t := *mmax - *mmin
    maxt := *mmax
    mint := *mmin

    left := dfs(n.Left, mmin, mmax)

    *mmax = maxt
    *mmin = mint
    right:= dfs(n.Right, mmin, mmax)

    return max(max(left, right), t)
}


func max(a , b int) int {
    if a > b {
        return a
    }

    return b
}

func min(a , b int) int {
    if a > b {
        return b
    }

    return a
}