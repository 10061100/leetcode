package main


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
    cache := make([]int, 101)
    maxLevel := 0

    dfs(root, 0, cache, &maxLevel)

    for i := 0; i < 101; i++ {
        if cache[i] > cache[maxLevel] {
            maxLevel = i
        }
    }

    return maxLevel
}


func dfs(n *TreeNode, deep int, cache []int, maxLevel *int) {

    if n == nil {
        return
    }

    (cache)[deep] += n.Val

    dfs(n.Left, deep+1, cache, maxLevel)
    dfs(n.Right, deep+1, cache, maxLevel)
}



