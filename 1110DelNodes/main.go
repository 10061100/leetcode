package main


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {

    m := make(map[int]bool)

    for _, v := range to_delete {
        m[v] = true
    }

    res := make([]*TreeNode, 0)


    delNodeSub(root, m, &res, true)

    return res
}


func delNodeSub(root *TreeNode, m map[int]bool, res *[]*TreeNode, first bool) {

    if root == nil {
        return
    }

    left := root.Left
    right:= root.Right

    // 进行分割
    if left != nil && m[left.Val] == true {
        root.Left = nil
    }

    if right != nil && m[right.Val] == true {
        root.Right = nil
    }

    isfirst := false

    if m[root.Val] == true {
        isfirst = true
    } else if first {
        *res = append(*res, root)
    }

    delNodeSub(left, m, res, isfirst)
    delNodeSub(right, m, res, isfirst)
}
