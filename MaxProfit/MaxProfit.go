// package MaxProfit

package main

import "fmt"

func main() {
    println(coinChange([]int{1}, 2));
}




func maxProfit(prices []int) int {

    if len(prices) < 2 {
        return 0
    }

    // 定制当前的状态的最大值
    has0No := 0 // 什么都没有, 当前呆着不动
    has0Bu := -prices[0] // 什么都没有, 然后买一个
    has1No := -prices[0] // 手上有一只股票, 然后什么都不做
    has1Sel:= 0  // 手上什么都没有, 然后买了一只股票

    for i := 1; i < len(prices); i ++ {

        // println(has0No, has0Bu, has1No, has1Sel)
        // 自己手里啥都不没有, 这次买了一个, 只能是 上次啥也没有
        o0Bu   := has0Bu
        has0Bu =  has0No - prices[i]

        // 自己这里什么都没有, 带着不动, 上次也啥都没有 or 上次有一个买了
        has0No = max(has0No, has1Sel)

        // 自己手里有一个, 什么都没做, 上次刚买 or 上次手里也有一个没动
        o1No   := has1No
        has1No = max(o0Bu, has1No)

        // 自己有利有一个买了, 则可能是: 上次刚买 or 上次手里有一个没动
        has1Sel = max(o1No + prices[i], o0Bu + prices[i])
    }
    // println(has0No, has0Bu, has1No, has1Sel)
    return max(has0No, has1Sel)
}


func maxProfit714(prices []int, fee int) int {
    if len(prices) < 2 {
        return 0
    }

    profit := 0
    hold   := -prices[0]

    for i := 1; i < len(prices); i++ {
        oldProfit := profit
        profit = max(profit, hold + prices[i] - fee)
        hold   = max(oldProfit - prices[i], hold)
    }

    return profit
}

func max( a , b int) int {
    if a > b {
        return a
    }

    return b
}



func coinChange(coins []int, amount int) int {
    if amount <= 0 {
        return 0
    }

    if len(coins) <= 0 {
        return -1
    }

    r := make([]int, amount+1)
    r[0] = 0

    for i  := 0; i < len(coins); i++ {
        if coins[i] <= amount {
            r[coins[i]] = 1
        }
    }

    for a := 1 ; a <= amount; a++ {
        for i := 0; i < len(coins); i++ {
            if coins[i] <= a {
                if a - coins[i] != 0 && r[a-coins[i]] != 0 {
                    if r[a] == 0 {
                        r[a] = r[a-coins[i]] + 1
                    } else {
                        r[a] = min(r[a], r[a-coins[i]] + 1)
                    }
                }
            }
        }
    }

    fmt.Println(r)

    if r[amount] > 0 {
        return r[amount]
    }

    return -1
}

func min(a, b int) int {
    if a > b {
        return b
    }

    return a
}

