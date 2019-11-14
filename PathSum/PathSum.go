// PathSum
package main

func main() {
    println(pathSum(nil, 1))
}

func pathSum(root *TreeNode, sum int) int {
    n , _ := onePath(root, sum)

    return n
}


func onePath(root *TreeNode, sum int) (int, map[int]int) {

    if nil == root {
        return 0, nil
    }

    lcount, lmap := onePath(root.Left, sum)
    rcount, rmap := onePath(root.Right, sum)

    summap := make(map[int]int)

    summap[root.Val] = 1
    if nil != lmap {
        for k, v := range lmap {
            nk := k + root.Val
            if _, ok := summap[nk]; !ok {
                summap[nk] = v
            } else {
                summap[nk] = summap[nk] + v
            }
        }
    }

    if nil != rmap {
        for k, v := range rmap {
            nk := k + root.Val
            if _, ok := summap[nk]; !ok {
                summap[nk] = v
            } else {
                summap[nk] = summap[nk] + v
            }
        }
    }

    tcount := 0
    for k, v := range summap {
        if k == sum {
            tcount += v
        }
    }

    return lcount + rcount + tcount, summap
}



type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}