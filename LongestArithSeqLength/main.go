package main

func main()  {
    println(longestArithSeqLength([]int{9,4,7,2,10}) == 3)
    println(longestArithSeqLength([]int{20,1,15,3,10,5,8}) == 4)
}

func longestArithSeqLength(A []int) int {
    l := len(A)

    if l <= 2 {
        return l
    }

    dp := make([]map[int]int, l)
    for i := 0; i < l; i++ {
        dp[i] = make(map[int]int)
        dp[i][0] = 1
    }

    max:= 0
    for i := 1; i < l; i++ {
        for j := 0; j < i; j++ {
            t := A[i] - A[j]
            dp[i][t] = GetMax(dp[i][t], dp[j][t] + 1)
            max = GetMax(dp[i][t], max)
        }
    }

    // fmt.Println(dp)
    return max+1
}


func GetMax(a , b int) int {
    if a > b {
        return a
    }

    return b
}
