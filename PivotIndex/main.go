// PivotIndex
package main

func main() {
    println(pivotIndex([]int{1, 7, 3, 6, 5, 6}) == 3)
    println(pivotIndex([]int{1, 2, 3}) == -1)
}

func pivotIndex(nums []int) int {


    sum := 0
    nlen:= len(nums)
    if nlen == 0 {
        return -1
    }

    for i := 0; i < nlen; i++ {
        sum += nums[i]
    }

    lsum := 0
    for i := 0; i < nlen; i++ {
        rsum := sum - nums[i] - lsum
        if rsum == lsum {
            return i
        }
        lsum += nums[i]
    }

    return -1
}

