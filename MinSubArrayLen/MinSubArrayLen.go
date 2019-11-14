// MinSubArrayLen
package main

func main() {
    println(minSubArrayLen(7, []int{2,3,1,2,4,3}))
}


func minSubArrayLen(s int, nums []int) int {
    t := 0
    l := len(nums) + 1
    start, end := 0, 0
    for ; end < len(nums);  end++ {
        t += nums[end]
        if t > s {
            // start挪到第一个 t > s的地方
            for ; start < end ; {
                if t - nums[start] >= s {
                    t -= nums[start]
                    start++
                } else {
                    break
                }
            }

            if l > end - start + 1 {
                l = end - start + 1
            }
        }
    }

    if l > len(nums) {
        return 0
    }

    return l
}