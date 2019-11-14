package main


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
    // res := make([]int, 0)

    i := k
    return midTravel(root, &i)
}


func midTravel(root *TreeNode, i *int) int {

    if root == nil {
        return -1
    }

    if t := midTravel(root.Left, i); t != -1 {
        return t
    }
    (*i)--
    if (*i) <= 0 {
        return root.Val
    }
    if t := midTravel(root.Right, i); t != -1 {
        return t
    }

    return -1
}
