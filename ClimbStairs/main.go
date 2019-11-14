
// climbStairs
package main

func main()  {
    println(climbStairs(1000))
}


func climbStairs(n int) int {
    if n == 0 || n == 1 {
        return 1
    }

    if n == 2 {
        return 2
    }

    step := make([]int, n + 1)
    step[0] = 1
    step[1] = 1
    step[2] = 2
    for i := 3; i < n + 1; i++ {
        step[i] = step[i-1] + step[i-2]
    }

    return step[n]
}
