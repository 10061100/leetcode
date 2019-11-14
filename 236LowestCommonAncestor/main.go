package main


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

var cache = make(map[*TreeNode]int)
var ans *TreeNode = nil

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }

    if root == p || root == q {
        return root
    }

    inTree(root, p, q)
    return ans
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }

    if root == p || root == q {
        return root
    }

    tc := inTree(root.Left, p, q)
    if tc == 1 {
        return root
    }

    if tc == 2 {
        return lowestCommonAncestor(root.Left, p, q)
    }

    return lowestCommonAncestor(root.Right, p, q)
}


func inTree(root, p, q *TreeNode) int {

    if root == nil {
        return 0
    }

    // if v, ok := cache[root]; ok {
    //     return v
    // }

    tc := 0
    if root == p || root == q {
        // ans = root
        tc = 1
    }

    leftc  := inTree(root.Left, p, q)

    if leftc + tc == 2 {
        if tc == 1 {
            ans = root
        }
        return 2
    }


    rightc := inTree(root.Right, p, q)
    if rightc + tc  == 2 {
        if tc == 1 {
            ans = root
        }
        return 2
    }

    if rightc + leftc + tc == 2 {
        ans = root
        return 2
    }

    return rightc + leftc + tc
    // cache[root] = tc + leftc + rightc

    // return tc + leftc + rightc
}