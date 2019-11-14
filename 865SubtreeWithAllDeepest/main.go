package main



type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }

    max := 0
    var anser *TreeNode = nil
    dfs(root, 1, &max, &anser)

    return anser
}


func dfs(node *TreeNode, deepth int, max *int, anwser **TreeNode) (length int) {
    if deepth > *max {
        *max = deepth
    }

    if node.Left == nil && node.Right == nil {
        if deepth == *max {
            *anwser = node
        }

        return deepth
    }


    // 判断左子树的深度和右子树的深度
    // 如果两个子树深度相等, 且等于max, 则就选这个点
    // 如果两个子树深度不等, 则选择较深的哪一个返回( =max )
    llen, rlen := deepth, deepth
    if node.Left != nil {
        llen = dfs(node.Left, deepth+1, max, anwser)
    }

    if node.Right != nil {
        rlen = dfs(node.Right, deepth+1, max, anwser)
    }

    if llen > rlen {
        length = llen
    } else {
        length = rlen
    }

    if llen == rlen  && llen >= *max {
        *anwser = node
    }

    return
}