// NextPermutation
package main

import "fmt"

func main() {
    a := []int{1}

    fmt.Println(":", search(a, 2))
}

func nextPermutation(nums []int)  {

    // 从最大的位置寻找一个降序节点( 即 n[i-1] < n[i] )
    start := len(nums) - 1
    for ; start >= 1 && nums[start] <= nums[start-1]; start-- {}

    if start > 0 {
        end := start
        // 如果有起始节点, 就开始找最小的笔起始大的节点
        for ; end < len(nums) - 1 && nums[start-1] <= nums[end+1]; end++ {}
        t := nums[start-1]
        nums[start-1] = nums[end]
        nums[end] = t
        fmt.Println(":", nums, ":", end)
    }

    // 开始从start开始到len(nums)-1进行位置调换
    for l, r  := start, len(nums) - 1; l <= r ; {
        if l != r {
            t := nums[l]
            nums[l] = nums[r]
            nums[r] = t
        }
        l++
        r--
    }

}


func search(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }

    if target == nums[len(nums) -1]{
        return len(nums) -1
    }

    for start, end := 0, len(nums) - 1; start < len(nums) && start <= end; {
        mid := start + (end - start)/2

        if nums[mid] == target {
            return mid
        } else if nums[start] <= nums[end] {
            // 一条线上
            if nums[mid] > target {
                end = mid-1
            } else {
                start = mid + 1
            }
        } else {
            if nums[mid] >= nums[start] {
                // mid start在一条线上
                if nums[mid] < target {
                    start = mid + 1
                } else {
                    // mid > target
                    if target >= nums[start] {
                        end = mid-1
                    } else {
                        start = mid + 1
                    }
                }
            } else {
                // mid end 在一条线上
                if nums[mid] > target {
                    end = mid-1
                } else {
                    if target >= nums[start] {
                        end = mid-1
                    } else {
                        start = mid + 1
                    }
                }
            }
        }
    }

    return -1
}
