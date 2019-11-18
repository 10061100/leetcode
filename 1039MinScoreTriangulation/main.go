package main


func minScoreTriangulation(A []int) int {

    n := len(A)

    if n == 3 {
        return A[0] * A[1] * A[2]
    }

    dp := make([][]int, n)
    for i := 0; i < n ; i++ {
        dp[i] = make([]int , n)
    }

    for l := 2; l < n; l++ {
        for i := 0; i < n; i++ {

        }
    }


}

