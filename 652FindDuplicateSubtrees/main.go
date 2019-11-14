package main

import (
    "strconv"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
    m := make(map[string][]*TreeNode)
    dfs(root, m)

    res := make([]*TreeNode, 0)
    for _, v := range m {
        if len(v) > 1 {
            res = append(res, v[0])
        }
    }

    return res
}


func dfs(root *TreeNode, m map[string][]*TreeNode) string{

    if root == nil {
        return ""
    }

    left := "=>"
    if root.Left != nil {
        left = dfs(root.Left, m)
    }

    right := "<="
    if root.Right != nil {
        right = dfs(root.Right, m)
    }

    hash := left + "|" +strconv.Itoa(root.Val) + "|" +right

    if _, ok := m[hash]; !ok {
        m[hash] = make([]*TreeNode, 0)
    }

    m[hash] = append(m[hash], root)

    return hash
}
