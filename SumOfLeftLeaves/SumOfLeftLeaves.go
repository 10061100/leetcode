// SumOfLeftLeaves
package main


// import "math"
import "fmt"

func main() {

    b := 0x7fff

    fmt.Printf("%d", -1 & b )
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
    sum := 0

    getSum(root, false, &sum)

    return sum
}


func getSum( root *TreeNode, isLeft bool, sum *int) {
    if root == nil {
        return
    }

    if isLeft {
        *sum += root.Val
    }

    getSum(root.Left, true, sum)
    getSum(root.Left, false, sum)
}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}