package main





func maxCoins(nums []int) int {
    l := len(nums)
    if l == 0 {
        return 0
    }

    if l == 1 {
        return nums[0]
    }

    dp := make([][]int, l)
    for i := 0; i < l; i++ {
        dp[i] = make([]int, l)
    }

    for i := 0; i < l ; i++ {
        dp[i][i] = nums[i] * getItem(nums, l, i-1) * getItem(nums, l, i+1)
    }

    // 子序列长度
    for i := 1; i <= l; i++ {
        // 对于每个子序列长度, 都从第一位开始遍历
        // 其中每个序列, 包括头尾
        for start := 0; start + i <= l; start++ {
            end := start + i - 1
            for j := start; j <= end; j++ {
                coins := nums[j] *  getItem(nums, l, start-1) * getItem(nums, l, end+1)
                if j > start {
                    coins += dp[start][j-1]
                }

                if j < end {
                    coins += dp[j+1][end]
                }

                if coins > dp[start][end] {
                    dp[start][end] = coins
                }
            }
        }

    }

    return dp[0][l-1]

}


func getItem(nums []int, l int, index int) int{
    if index >= l || index < 0 {
        return 1
    }

    return nums[index]
}
