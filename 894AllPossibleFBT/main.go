package main



type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


func allPossibleFBT(N int) []*TreeNode {
    res := make([]*TreeNode, 0)

    var root, parent *TreeNode
    dfs(N, false, parent, &root, true, &res)

    return res
}


func NewTreeNode() *TreeNode {
    return &TreeNode{
        Val: 0,
        Left: nil,
        Right: nil,
    }
}

func dfs(n int, isRight bool, parent *TreeNode, root **TreeNode, islast bool, res *[]*TreeNode) {

    if n % 2 == 0 {
        return
    }

    t := NewTreeNode()

    if root == nil {
        *root = t
    }

    if parent != nil {
        if isRight {
            parent.Right = t
        } else {
            parent.Left = t
        }
    }

    if n != 1 {
        n--
        for left := 1; left < n; left += 2 {
            right := n - left
            if right % 2 == 0 {
                continue
            }

            dfs(left, false, t , root, false, res)
            dfs(left, true, t, root, islast, res)
        }
    } else {
        if islast {
            // copy root
            *res = append(*res, copyTree(*root))
        }
    }

}


func copyTree(root *TreeNode) *TreeNode {

    if root == nil {
        return nil
    }

    t := NewTreeNode()

    t.Left = copyTree(root.Left)
    t.Right = copyTree(root.Right)

    return t
}