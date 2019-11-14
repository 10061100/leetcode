package main

func main() {
    println(shortestCommonSupersequence("ab", "bc"))
    println(shortestCommonSupersequence("abac", "cab"))
}


func shortestCommonSupersequence2(str1 string, str2 string) string {
    if len(str2) > len(str1) {
        str2, str1 = str1, str2 // str1 代表比较长的字符串
    }

    // 判断str2在
    dp := make([][]string, len(str2)+1)
    for i := 0 ; i <= len(str2); i++ {
        dp[i] = make([]string, len(str1) + 1)
    }

    dp[0][0] = ""
    for i := 0; i < len(str1); i++ {
        dp[0][i+1] = str1[0:i+1]
    }

    for i := 0; i < len(str2); i++ {
        dp[i+1][0] = str2[0:i+1]
    }

    for i := 1; i <= len(str2); i++ {
        for j := 1; j <= len(str1); j++{
            if str1[j-1] == str2[i-1] {
                dp[i][j] = dp[i-1][j-1] + str1[j-1:j]
            } else {
                dp[i][j] = minstr(dp[i-1][j] + str2[i-1:i], dp[i][j-1] + str1[j-1:j])
            }
        }
    }

    return dp[len(str2)][len(str1)]
}


func shortestCommonSupersequence(str1 string, str2 string) string {
    if len(str2) > len(str1) {
        str2, str1 = str1, str2 // str1 代表比较长的字符串
    }

    // 判断str2在
    dp := make([]string, len(str1)+1)
    tdp:= make([]string, len(str1)+1)

    dp[0] = ""
    for i := 0; i < len(str1); i++ {
        dp[i+1] = str1[0:i+1]
    }

    copy(tdp, dp)

    for i := 1; i <= len(str2); i++ {
        dp[0] = str2[0:i]
        for j := 1; j <= len(str1); j++{
            if str1[j-1] == str2[i-1] {
                dp[j] = tdp[j-1] + str1[j-1:j]
            } else {
                dp[j] = minstr(dp[j] + str2[i-1:i], dp[j-1] + str1[j-1:j])
            }
        }
        copy(tdp, dp)
        // fmt.Println(dp)
    }

    return dp[len(str1)]
}


func minstr( a, b string) string {
    if len(a) > len(b) {
        return b
    }

    return a
}
