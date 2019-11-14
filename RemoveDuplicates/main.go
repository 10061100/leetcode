// RemoveDuplicates
package main

import "fmt"

func main()  {
    array := []int{1,1,1}
    println(removeDuplicates(array))
    fmt.Println(array)
}

func removeDuplicates(nums []int) int {
    if len(nums) <= 1 {
        return len(nums)
    }

    length := 1
    cur := nums[0]
    lastIdx := len(nums)

    for i := 0; i < lastIdx;  {
        // 找到第一个不一样的值
        j := i
        for ; j < lastIdx && nums[j] == cur; j++ {}
        if j >= lastIdx {
            break // 直接退出
        }

        nums[length] = nums[j]
        length++
        cur = nums[j]

        i = j

    }

    nums = nums[0:length]

    return length
}
