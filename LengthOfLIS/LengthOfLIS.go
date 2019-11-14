package main

func main() {
    print(lengthOfLIS([]int{10,9,2,5,3,7,101,18}) == 4)
}


func lengthOfLIS(nums []int) int {

    des := make([]int, len(nums))

    des[0] = 1
    for i := 1; i < len(nums); i++ {
        des[i] = 1;
        for j := 0; j < i; j++ {
            if nums[j] < nums[i] && des[i] < des[j] + 1 {
                des[i] = des[j] + 1
            }
        }
    }

    max := 1
    for i := 0 ; i < len(nums); i++ {
        if max < des[i] {
            max = des[i]
        }
    }

    return max
}