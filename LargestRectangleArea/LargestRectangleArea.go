// LargestRectangleArea
package main

import (
    // "fmt"
)

func main() {
    firstMissingPositive([]int{1,2,0})
}

func largestRectangleArea(heights []int) int {
    l := len(heights)
    if l == 0 {
        return 0
    }

    stack := make([]int, l)
    top := 0
    max := 0
    heights = append(heights, 0)
    heights[l] = 0
    for i := 0; i <= l; i++ {
        if top == 0 || heights[stack[top-1]] <= heights[i] {
            stack[top] = i
            top++
        } else {
            t := stack[top-1]
            top--

            area := 0
            if top == 0 {
                area = heights[t] * i
            } else {
                area = heights[t] * (i - stack[top-1] - 1)
            }

            if max < area {
                max = area
            }
            i--
        }
    }

    return max
}


func firstMissingPositive(nums []int) int {
    l := len(nums)

    if l <= 0 {
        return 1
    }

    for i := 0 ; i < l; i++ {
        println(i)
        for ; nums[i] <= l && nums[i] >= 1 && nums[nums[i]-1] != nums[i]; {
            println(i)
            t := nums[nums[i]-1]
            nums[nums[i]-1] = nums[i]
            nums[i] = t
        }
    }

    for i := 0; i < l; i++ {
        if nums[i] != i+1 {
            return i+1
        }
    }

    return l+1
}







