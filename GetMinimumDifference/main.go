
// GetMinimumDifference
package main

func main() {

}

func getMinimumDifference(root *TreeNode) int {

    pre, depth := -1, -1

    dfs(root, &pre, &depth)

    return depth
}


func dfs(node *TreeNode, pre *int, depth *int){
    if nil != node.Left {
        dfs(node.Left, pre, depth)
    }

    if *pre != -1 {
        if *depth == -1 || *depth > node.Val - *pre {
            *depth = node.Val - *pre
        }
    }
    *pre = node.Val

    if nil != node.Right {
        dfs(node.Right, pre, depth)
    }
}



func hdfs(node *TreeNode, pre *int){
    if nil != node.Right {
        hdfs(node.Right, pre)
    }

    t := node.Val
    node.Val += *pre
    *pre = t

    if nil != node.Left {
        hdfs(node.Left, pre)
    }
}



type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}