package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}



func constructMaximumBinaryTree(nums []int) *TreeNode {
    return rec(nums, 0, len(nums)-1)
}


func rec(nums []int, start, end int) *TreeNode {
    if start > end {
        return nil
    }

    maxIdx := start
    for i := start + 1; i <= end; i++ {
        if nums[i] > nums[maxIdx] {
            maxIdx = i
        }
    }

    node := NewTreeNode(nums[maxIdx])

    node.Left = rec(nums, start, maxIdx-1)
    node.Left = rec(nums, maxIdx+1, end)

    return node
}


func NewTreeNode(v int) *TreeNode {
    return &TreeNode{
        Val: v,
        Left: nil,
        Right: nil,
    }
}
