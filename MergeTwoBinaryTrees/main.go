// MergeTwoBinaryTrees
package main

func main() {

}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    if t1 == nil {
        return t2
    }

    mergeTwoBinaryTrees(t1, t2)

    return t1
}


func mergeTwoBinaryTrees(t1 *TreeNode, t2 *TreeNode) {
    if t1 == nil {
        return
    }

    if t2 == nil {
        return
    }

    t1.Val += t2.Val
    if t1.Left == nil {
        t1.Left = t2.Left
        t2.Left = nil
    }

    if t1.Right == nil {
        t1.Right == t2.Right
        t2.Right == nil
    }

    mergeTwoBinaryTrees(t1.Left, t2.Left)
    mergeTwoBinaryTrees(t1.Right, t2.Right)
}