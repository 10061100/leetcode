package main



type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }

    return buildBTree(preorder, inorder, 0, len(preorder)-1, 0, len(preorder)-1)
}

func buildBTree(preorder []int, inorder []int, prestart, preend, instart, inend int) *TreeNode {

    if prestart > preend {
        return nil
    }

    if preend == prestart {
        return NewTree(preorder[prestart])
    }

    cur := preorder[prestart]

    mid := instart
    for ; mid <= inend && inorder[mid] != cur; mid ++ {}

    // 找到根节点
    root := NewTree(cur)

    // 左子树为mid左边的, 右子树为右边的
    leftcount := (mid - 1) - instart + 1
    rightcount:= inend - (mid + 1) + 1
    left := buildBTree(preorder, inorder, prestart+1,prestart+leftcount, instart, mid-1)
    right := buildBTree(preorder, inorder, preend-rightcount+1, preend, mid+1, inend)

    root.Left = left
    root.Right = right

    return root
}


func NewTree(v int) *TreeNode {
    return &TreeNode{
        Val:v,
        Left: nil,
        Right: nil,
    }
}
