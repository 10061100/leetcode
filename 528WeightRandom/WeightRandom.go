package main

type Solution struct {
    dp []int
    last int
}


func Constructor(w []int) Solution {
    s := Solution{
        dp: make([]int, len(w)),
    }

    s.dp[0] = w[0]
    for i := 1; i < len(w); i++ {
        s.dp[i] = s.dp[i-1] + w[i]
    }

    s.last = s.dp[len(s.dp)-1]
    return s
}


func (this *Solution) PickIndex() int {
    // location := rand.Intn(this.last)

    // sort.SearchInts()
    return 1
}
