// FindShortestSubArray
package main

func main() {
    println(findShortestSubArray([]int{1,2,2,3,1,4,2})  == 6 )
}

const COUNT = 0
const LEFT  = 1
const RIGHT = 2

func findShortestSubArray(nums []int) int {
    hash := make(map[int][]int)

    length := len(nums)
    min := length
    degree := 0
    for i := 0; i < length; i++ {
        res, ok := hash[nums[i]]
        if !ok {
            res = make([]int, 3)
            res[COUNT] = 0
            res[LEFT]  = i
            res[RIGHT] = 0
        }

        res[RIGHT] = i
        res[COUNT]++
        if res[COUNT] >= degree {
            diff := res[RIGHT] - res[LEFT] + 1;

            if res[COUNT] > degree {
                min = diff
            } else {
                // 如果相等话,取较小的
                if diff < min {
                    min = diff
                }
            }

            degree = res[COUNT]
        }

        hash[nums[i]] = res
    }

    return min
}