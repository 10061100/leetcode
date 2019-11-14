// JudgeSquareSum
package main

func main() {
    println(judgeSquareSum(9) == true)
    println(judgeSquareSum(1) == true)
    println(judgeSquareSum(100) == true)
    println(judgeSquareSum(4) == true)
    println(judgeSquareSum(101) == true)
}

func judgeSquareSum(c int) bool {
    for i := 0 ; i * i <= c ; i++ {
        if v := c - i*i; binarySearchIs2(v) {
            return true
        }
    }

    return false
}


func binarySearchIs2( n int) bool {
    min := 0
    max := n

    if min * min == n {
        return true
    }

    if max * max == n {
        return true
    }

    for ;; {

        cur := min + (max - min)/2

        if cur * cur == n {
            return true
        }

        if max <= min {
            return false
        }

        if cur * cur < n {
            min = cur + 1
        } else {
            max = cur
        }
    }
}