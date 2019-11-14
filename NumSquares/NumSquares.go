package main

func main() {
    println(minimumDeleteSum("seat", "eaty"))
}

func minimumDeleteSum(s1 string, s2 string) int {
    bs1 := []byte(s1)
    bs2 := []byte(s2)

    l1 := len(bs1)
    l2 := len(bs2)

    if l1 == 0 && l2 == 0 {
        return 0
    }

    dp := make([][]int, l1+1)
    for i := 0; i <= l1; i++ {
        dp[i] = make([]int, l2+1)
    }

    dp[0][0] = 0
    for i := 1; i <= l2; i++{
        dp[0][i] = dp[0][i-1] + getCode(bs2[i-1])
    }

    for i := 1; i <= l1; i++ {
        dp[i][0] = dp[i-1][0] + getCode(bs1[i-1])
    }

    for i := 1; i <= l1; i++ {
        for j := 1; j <= l2; j++ {
            if bs1[i-1] == bs2[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                ds1 := dp[i-1][j] + getCode(bs1[i-1])
                ds2 := dp[i][j-1] + getCode(bs2[j-1])
                if ds1 < ds2 {
                    dp[i][j] = ds1
                } else {
                    dp[i][j] = ds2
                }
            }
        }
    }

    return dp[l1][l2]
}


func getCode(c byte) int {
    return int(c - 'a')
}
