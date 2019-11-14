// MaximumProduct
package main

func main() {
    println(maximumProduct([]int{-1,-1,0}) == 0)
    println(maximumProduct([]int{-1,0,0}) == 0)
    println(maximumProduct([]int{-1,-1,-1}) == -1)
}


func maximumProduct(nums []int) int {

    min1, min2 := 1001, 1001
    max1, max2, max3 := -1001, -1001, -1001

    for _, v := range nums {
        if v < min1 {
            min2 = min1
            min1 = v
        } else if v < min2 {
            min2 = v
        }

        if v > max3 {
            max1 = max2
            max2 = max3
            max3 = v
        } else if v > max2 {
            max1 = max2
            max2 = v
        } else if v > max1 {
            max1 = v
        }
    }

    v1 := max1 * max2 * max3
    v2 := min1 * min2 * max3

    if v1 > v2 {
        return v1
    }

    return v2
}

