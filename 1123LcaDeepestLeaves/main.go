package main


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
    _, res := dfs(root, 1)

    return res
}




func dfs(node *TreeNode, cdeep int) (deep int, res *TreeNode) {

    if node == nil {
        return -1, nil
    }

    if node.Left == nil && node.Right == nil {
        return cdeep, node
    }


    ldeep, lres := dfs(node.Left, cdeep + 1)
    rdeep, rres := dfs(node.Right,cdeep + 1)

    if ldeep == rdeep {
        return ldeep, node
    }

    if ldeep > rdeep {
        return ldeep, lres
    }

    return rdeep, rres
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
