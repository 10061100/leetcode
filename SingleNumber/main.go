package main

func main() {
    println(singleNumber([]int{1,1,2}))
}

func singleNumber(nums []int) int {
    l := len(nums)

    if l == 1 {
        return nums[0]
    }

    ret := nums[0]
    for i := 1; i < l; i++ {
        // fmt.Println(ret, nums[i], ^nums[i])
        ret = ret^nums[i]
    }


    return ret
}
