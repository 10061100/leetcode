package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func main()  {
    
}


func deleteNode1(root *TreeNode, key int) *TreeNode {

    des, parent := findParent(root, key)
    // 没找到对应节点, 不用删除
    if des == nil {
        return root
    }


    // 找到了对应结点, 需要用它右子树的最小值来替换当前的树
    if des.Right == nil {
        // 如果是根节点的话
        if parent == nil {
            return des.Left
        }

        // 如果不是根节点, 就用右子树替换自己
        if parent.Left == des {
            parent.Left = des.Left
        } else {
            parent.Right = des.Left
        }

        return root
    }

    // 如果右子树不是空的, 就取右子树最小的值, 替换掉当前的值
    maxNode, maxParent := getTreeMax(des.Right, des)
    newNode := maxNode
    newRight:= newNode.Right

    newNode.Left = des.Left
    if maxParent != des {
        newNode.Right = des.Right
        maxParent.Left = newRight
    }

    if parent == nil {
        root = newNode
    } else {
        if parent.Left == des {
            parent.Left = newNode
        } else {
            parent.Right = newNode
        }
    }

    return root
}

func getTreeMax(node *TreeNode, parent *TreeNode) (des *TreeNode, newParent *TreeNode) {
    if node == nil {
        return nil, parent
    }

    for ; node != nil; {
        if node.Left == nil {
            return node, parent
        }

        parent = node
        node = node.Left
    }

    return node, parent
}


func findParent(root *TreeNode, key int) (cur *TreeNode, parent *TreeNode) {
    parent = nil
    cur = root
    for ; cur != nil ; {
        if cur.Val == key {
            return cur, parent
        } else if cur.Val > key {
            parent = cur
            cur = cur.Left
        } else {
            parent = cur
            cur = cur.Right
        }
    }

    return nil, nil
}


func deleteNode(root *TreeNode, key int) *TreeNode {
    des, parent := findParent(root, key)

    if des == nil {
        return nil
    }

    newTree := del(des)

    if parent == nil {
        return newTree
    }

    isLeft := true
    if parent.Right == des {
        isLeft = false
    }

    if isLeft {
        parent.Left = newTree
    } else {
        parent.Right = newTree
    }

    return root
}


func del(node *TreeNode) *TreeNode {

    if node == nil {
        return node
    }

    // 删除当前的点, 首要的删除自己的左子树
    if node.Right == nil {
        return node.Left
    }

    // 如果右子树, 就删除右子树的最左子树
    minNode, minParent := getTreeMax(node.Right, node)
    isLeft := true
    if minParent.Right == minNode {
        isLeft = false
    }

    newTree := del(minNode)
    if isLeft {
        minParent.Left = newTree
    } else {
        minParent.Right = newTree
    }

    minNode.Left = node.Left
    minNode.Right = node.Right

    return minNode
}