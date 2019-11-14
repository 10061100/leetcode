package main


func findClosestElements(arr []int, k int, x int) []int {

    size := len(arr)
    left := 0
    right:= size - k

    for ; left < right ; {
        mid := (left + right) >> 1

        min := arr[mid]
        max := arr[mid + k]


        if max - x >= x - min {
            right = mid
        } else {
            left = mid + 1
        }
    }


    return arr[left:left+k]
}
