// CombinationSum4
package main

func main() {
    println(combinationSum4([]int{1, 2, 3}, 2000))
}


func combinationSum4(nums []int, target int) int {
    cache := make(map[int]int)
    if target <= 0 {
        return target
    }

    return count(nums, target, cache)
}


func  count(nums []int, target int, cache map[int]int) int {
    if t , ok := cache[target]; ok {
        return t
    }

    t := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == target {
            t++
        } else if nums[i] > target {
            continue
        } else {
            t  = t + count(nums, target - nums[i], cache)
        }
    }

    cache[target] = t

    return t
}