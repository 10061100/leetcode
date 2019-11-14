package main


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


func main() {

}

func distributeCoins(root *TreeNode) int {

    if root == nil {
        return 0
    }

    _, v := subCoins(root)

    return v
}


func subCoins(root *TreeNode) (overload, up int) {

    if root == nil {
        return 0, 0
    }

    leftoverload, leftup := subCoins(root.Left)
    rightoverload,rightup:= subCoins(root.Right)

    overload = root.Val + leftoverload + rightoverload - 1
    // steps    = leftsteps + rightsteps
    up       = leftup + rightup + abs(overload)// 还需要包括剩下的过载

    return
}

func abs(a int) int {
    if a > 0 {
        return a
    }

    return -a
}
