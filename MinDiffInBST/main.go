// MinDiffInBST
package main


func main() {
    println(minDiffInBST(buildTrees([]int{90,94,96})))
}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

type Stack struct {
    Container []*TreeNode
    Top int
}

func NewStack() *Stack {
    c := []*TreeNode{}
    return &Stack{
        Container:c,
        Top:0,
    }
}

func (s *Stack) Push(l *TreeNode) {
    s.Container = append(s.Container, l)
    s.Top++
}

func (s *Stack) Pop() *TreeNode {
    if s.Top <= 0 {
        return nil
    }
    s.Top--
    tmp := s.Container[s.Top]
    s.Container = s.Container[:s.Top]

    return tmp
}

// 中序遍历
func minDiffInBST(root *TreeNode) int {
    s := NewStack()
    cur := root
    var befor *TreeNode
    min := -1
    for ;; {

        // 一直到最左边的值
        for ; cur != nil; cur = cur.Left {
            s.Push(cur)
        }
        if befor == nil {
            befor = s.Pop()
            if befor == nil {
                break
            }
            cur = befor.Right
            continue
        }

        t := s.Pop()
        if t == nil {
            break
        }
        if v := t.Val - befor.Val; min < 0 || min > v {
            min = v
        }

        befor = t
        cur = t.Right
    }

    return min
}

func buildTrees(a []int) *TreeNode {
    if len(a) == 0 {
        return nil
    }

    head := &TreeNode{
        Val:a[0],
        Left:nil,
        Right:nil,
    }

    for i := 1; i <  len(a); i++ {
        v := a[i]
        if v == -1 {
            continue
        }
        cur := head
        for ;; {
            if cur.Val > v {
                if cur.Left != nil {
                    cur = cur.Left
                } else {
                    cur.Left = &TreeNode{
                        Val:v,
                        Left:nil,
                        Right:nil,
                    }
                    break
                }
            } else {
                if cur.Right != nil {
                    cur = cur.Right
                } else {
                    cur.Right = &TreeNode{
                        Val:v,
                        Left:nil,
                        Right:nil,
                    }
                    break
                }
            }
        }
    }

    return head
}




