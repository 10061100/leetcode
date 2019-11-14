package main

import (
    "fmt"
)

func main() {
    fmt.Println(nthSuperUglyNumber(100000, []int{7,19,29,37,41,47,53,59,61,79,83,89,101,103,109,127,131,137,139,157,167,179,181,199,211,229,233,239,241,251}))
}

func nthSuperUglyNumber(n int, primes []int) int {

    if n == 1 {
        return 1
    }

    dp := make([]int, n)
    store := make([]int, len(primes))
    dp[0] = 1

    for i := 1; i < n; i++ {
        dp[i] = 8888888888
        for j := 0; j < len(primes); j++ {

            t := dp[store[j]] * primes[j]
            if t < dp[i] {
                dp[i] = t
            }
        }

        for j := 0; j < len(primes); j++ {
            if dp[store[j]] * primes[j] <= dp[i] {
                store[j]++
            }
        }
    }

    return dp[n-1]
}
