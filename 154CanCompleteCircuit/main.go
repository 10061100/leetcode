package main

func canCompleteCircuit(gas []int, cost []int) int {

    diff := make([]int, len(gas))
    sum := 0

    // diff表示在每个点, 如果想开到下一个站点, 油箱里面最少需要多少油, 负数代表可以透支
    for i := 0; i < len(gas); i++ {
        diff[i] = cost[i] - gas[i]
        sum += diff[i]
    }

    if sum > 0 {
        return -1
    }

    startNode := 0
    cur := 0 // 代表油箱里剩余了多少油
    sum = 0 // 代表总共需要多少油, 如果从零点出发的话, 最起码保证油箱里面有多少油
    for i := 0; i < len(gas); {
        sum += diff[i]
        cur += diff[i]
        // 如果这个点需要油, 那就能是起点
        if cur > 0 {
            startNode = i+1
            cur = 0
        }
    }

    return startNode
}
