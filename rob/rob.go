package main

func main() {

}




type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func rob(root *TreeNode) int {
    return robRoot(root, true)
}


// robRoot false = <robRoot.left, true> + <robRoot.right, true>
// robRoot true  = max(<roobRoot, false>, <robRoot.left, false> + <robRoot.right, false>)
func robRoot(root *TreeNode, rootRob bool) int {
    profit := 0
    if root == nil {
        return profit
    }

    profit = robRoot(root.Left, true) + robRoot(root.Right, true)
    if rootRob == false {
        // 如果当前的节点不能抢劫
        return profit
    }

    // 如果当前节点可以抢劫
    t := robRoot(root.Left, false) + robRoot(root.Right, false) + root.Val

    if  t > profit {
        return t
    }


    return profit
}