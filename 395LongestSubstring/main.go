package main

import (
    "fmt"
)

func main()  {
    fmt.Println(longestSubstring("aaabb", 4))
}

func longestSubstring(s string, k int) int {
    ss := []byte(s)

    return subJob(ss, 0, len(ss)-1, k)
}

func subJob(ss []byte, start int, end int, k int) int {

    t := [26]int{}

    for i := start; i <= end; i++ {
        t[ss[i] - 'a']++
    }

    for ; start < end; start++ {
        if t[ss[start] - 'a'] >= k {
            break
        }
    }

    for ; end > start; end-- {
        if t[ss[end] - 'a'] >= k {
            break
        }
    }

    if end <= start {
        return 0
    }

    for i := start + 1; i < end; i++ {
        if t[ss[i]-'a'] < k {
            return max(subJob(ss, start, i-1,k), subJob(ss, i+1, end,k))
        }
    }

    return end - start +1
}


func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}