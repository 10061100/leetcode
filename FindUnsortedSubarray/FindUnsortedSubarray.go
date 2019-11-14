// FindUnsortedSubarray
package main

func main() {
    println(findUnsortedSubarray([]int{1,3,2,2,2}) )
    println(findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}) )
}

func findUnsortedSubarray(nums []int) int {

    fir := 0
    las := len(nums) - 1
    nlen:= len(nums)

    for fir = 0; fir < nlen-1 ; fir++ {
        if nums[fir] > nums[fir+1]{
            break
        }
    }

    if fir == nlen-1 {
        return 0
    }

    for las = nlen - 1; las > 0; las-- {
        if nums[las] < nums[las-1] {
            break
        }
    }

    min := 2147483647
    max := -2147483648

    for i := fir ; i <= las ; i++ {
        if nums[i] > max {
            max = nums[i]
        }

        if nums[i] < min {
            min = nums[i]
        }
    }

    // println(max)
    // println(las)

    left := 0
    right:= 0

    for left = 0; left <= fir && nums[left] <= min; left++ {}
    for right = nlen-1; right >= las && nums[right] >= max; right-- {
        // println(right)
    }

    return right-left+1
}