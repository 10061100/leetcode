// MinCostClimbingStairs
package main

func main() {
    println(dominantIndex([]int{1,2,16,35,44,100,27,0}))
}


func minCostClimbingStairs(cost []int) int {
    sum := make([]int, len(cost)+1)
    sum[1-1] = 0
    sum[2-1] = 0
    for i := 3; i <= len(cost) + 1; i++ {
        step1 := sum[i-2-1] + cost[i-2-1]
        step2 := sum[i-1-1] + cost[i-1-1]
        if step1 > step2 {
            sum[i-1] = step2
        } else {
            sum[i-1] = step1
        }
    }

    return sum[len(cost)]
}


func minCostClimbingStairsAtK(cost []int , k int) int {
    if k == 0 || k == 1 || k < 0 {
        return 0
    }

    step1 := minCostClimbingStairsAtK(cost, k-1) + cost[k-1]
    step2 := minCostClimbingStairsAtK(cost, k-2) + cost[k-2]

    if step1 <= step2 {
        return step1
    }

    return step2
}





