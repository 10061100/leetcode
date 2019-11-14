
package main

import (
    "fmt"
)

func main() {

    // println(minCut("efe"))
    // a := []int{2,0,2,1,1,0}
    // sortColors(a)
    // fmt.Println(a)

    num1 := []int{3, 9}
    num2 := []int{8, 9}

    fmt.Println(maxNumber(num1, num2, 3))
}


func minCut(s string) int {
    bs := []byte(s)
    ls := len(bs)

    if ls == 0 || ls == 1 {
        return 0
    }

    dp := make([]int, ls+1)
    for i := 1; i <= ls; i++ {
        dp[i] = i-1
    }

    for i := 0 ; i < ls; i++ {
        // 奇数的回文
        for j := 0; i - j >= 0 && i + j < ls && bs[i-j] == bs[i+j]; j++ {
            // fmt.Println(dp[i+j+1], dp[i-j])
            now := 0
            if i - j == 0 {
                now = 0
            } else {
                now = dp[i-j] + 1
            }
            if dp[i+j+1] > now {
                dp[i+j+1] = now
            }
        }

        // 偶数的回文
        for j:= 1; i-j+1 >= 0 && i +j < ls && bs[i-j+1] == bs[i+j]; j++ {
            // fmt.Println(dp[i+j+1], dp[i-j+1])
            now := 0
            if i - j + 1 == 0 {
                now = 0
            } else {
                now = dp[i-j+1] + 1
            }
            if dp[i+j+1] > now {
                dp[i+j+1] = now
            }
        }
    }

    fmt.Println(dp)

    return dp[ls]
}


func getMoneyAmount(n int) int {
    if n == 1 {
        return 0
    }

    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, n)
    }

    dp[0][0] = 0
    for i := 2; i <= n; i++ {
        for start :=0; start + i <= n; start++ {
            end := start + i - 1
            dp[start][end] = (start+1) + dp[start+1][end]
            for j := start + 1; j <= end; j++ {
                left := dp[start][j-1] + j + 1
                right:= j + 1
                if j != end {
                    right += dp[j+1][end]
                }

                if ( left >= right ) && dp[start][end] > left {
                    dp[start][end] = left
                }

                if right > left && dp[start][end] > right {
                    dp[start][end] = right
                }
            }
        }

        fmt.Println(dp)
    }

    return dp[0][n-1]
}


func sortColors(nums []int)  {
    l := len(nums)

    if l <= 1 {
        return
    }

    zero := 0
    second:= l-1

    for i := 0; i < l; {
        if nums[i] == 0 && i > zero {
            nums[i] = nums[zero]
            nums[zero] = 0
            zero++
        } else if nums[i] == 2 && i < second {
            nums[i] = nums[second]
            nums[second] = 2
            second--
        } else {i++}

    }
}

func min( a , b int) int {
    if a > b {
        return b
    }

    return a
}

func maxNumber(nums1 []int, nums2 []int, k int) []int {
    m := len(nums1)
    n := len(nums2)
    rl:= min(m+n, k)
    r := make([]int, rl)
    if k == 0 || (m==0 && n==0) {
        return r
    }

    cache := make([][][]int, m+1)
    for i := 0; i <= m; i++ {
        cache[i] = make([][]int, n+1)
        for j := 0; j <= n; j++ {
            cache[i][j] = make([]int, k+1)
        }
    }

    max := 0
    for tk := 1; tk <= k; tk++ {
        // fmt.Print("{")

        for tm := 0; tm <= m ; tm++ {
            // fmt.Print("[")
            for tn := 0; tn <= n; tn++ {
                // fmt.Println(tm, tn, tk)
                // fmt.Println(cache)
                t1 , t2, t3, t4 := 0, 0, 0, 0
                if tm >= 1 {
                    t1 = cache[tm-1][tn][tk-1] * 10 + nums1[tm-1]
                    t3 = cache[tm-1][tn][tk]
                }

                if tn >= 1 {
                    t2 = cache[tm][tn-1][tk-1] * 10 + nums2[tn-1]
                    t4 = cache[tm][tn-1][tk]
                }

                if t1 >= t2 && t1 >= t3 && t1 >= t4 {
                    cache[tm][tn][tk] = t1
                }else if t2 >= t1 && t2 >= t3 && t2 >= t4  {
                    cache[tm][tn][tk] = t2
                }else if t3 >= t1 && t3 >= t2 && t3 >= t4 {
                    cache[tm][tn][tk] = t3
                }else {
                    cache[tm][tn][tk] = t4
                }

                if cache[tm][tn][tk] >= max {
                    max = cache[tm][tn][tk]
                }
                // fmt.Printf("%d ", cache[tm][tn][tk])
            }
            // fmt.Print("] ")
        }
        // for i := 0; i <= m; i++ {
        //     fmt.Print("[")
        //     for j := 0; j <= n; j++ {
        //         fmt.Printf("%d ", cache[i][j][tk])
        //     }
        //     fmt.Print("] ")
        // }
        // fmt.Println("}")
    }

    for i := rl-1; i >= 0; i-- {
        r[i] = max % 10
        max /= 10
    }

    // fmt.Println(max)
    return r
}
