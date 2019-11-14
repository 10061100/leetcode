package main

func main() {
    println(numTrees(3))
    // println(numBTree(1,1))
}

func numTrees(n int) int {
    // return numBTree(1, n)

    dp := make([]int, n+1)
    dp[0] = 1
    dp[1] = 1

    for i := 2; i <= n; i++ {
        for j := 1; j <= i; j++ {
            dp[i] += dp[i-j] * dp[j-1]
        }
    }
    // fmt.Println(dp)

    return dp[n]
}


func numBTree(from, to int) int {
    if from == to {
        return 1
    }

    if to - from == 1 {
        return 2
    }

    t := numBTree(from+1, to) + numBTree(from, to-1)
    for  i := from+1; i <= to-1; i++ {
        t += numBTree(from, i-1) * numBTree(i+1, to)
    }

    return t
}