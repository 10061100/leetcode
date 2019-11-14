// IntegerBreak
package main

func main() {
    println(integerBreak(8))
}

func integerBreak(n int) int {
    if n == 0 {
        return 0
    }

    if n == 1 {
        return 0
    }

    if n == 2 {
        return 1
    }

    if n == 3 {
        return 2
    }

    tmp := make([]int, n+1)
    tmp[0] = 0
    tmp[1] = 1
    tmp[2] = 2
    tmp[3] = 3

    for i := 4; i <= n; i++ {
        tmp[i] = i
        for j := 0; j <= (i+1)/2 ; j++ {
            if tmp[j] * tmp[i-j] > tmp[i] {
                tmp[i] = tmp[j] * tmp[i-j]
            }
        }
    }

    return tmp[n]
}
