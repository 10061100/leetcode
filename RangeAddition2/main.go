// RangeAddition2
package main

func main() {
    println(maxCount(3,3, [][]int{{2,2},{3,3}}) == 4)
}

func maxCount(m int, n int, ops [][]int) int {
    mmin := m
    nmin := n

    for i := 0; i < len(ops); i++ {
        if mmin > ops[i][0] {
            mmin = ops[i][0]
        }

        if nmin > ops[i][1] {
            nmin = ops[i][1]
        }
    }

    return mmin * nmin
}