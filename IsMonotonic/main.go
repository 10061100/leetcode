package main


func main() {
    println(isMonotonic([]int{6,5,4,4}))
    println(isMonotonic([]int{1,3,2}))
}

func isMonotonic(A []int) bool {
    if len(A) <= 2 {
        return true
    }

    dir := 0
    for i := 1; i < len(A); i++ {
        newdir := 0
        if A[i] > A[i-1] {
            newdir = 1
        } else if A[i] < A[i-1] {
            newdir = -1
        }

        if newdir == 0 {
            continue
        }

        if dir == 0 {
            dir = newdir
            continue
        }

        if dir != newdir {
            // println(i, A[i], A[i-1])
            // println(dir, newdir)
            return false
        }
    }

    return true
}