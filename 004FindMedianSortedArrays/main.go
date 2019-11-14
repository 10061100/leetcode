package main

import (
    "fmt"
)

func main() {
    fmt.Println(findMedianSortedArrays([]int{1,1,3}, []int{1}))
}


func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    n1 := len(nums1)
    n2 := len(nums2)

    if n1 > n2 {
        nums1, nums2 = nums2, nums1
        n1, n2 = n2, n1
    }

    left := (n1+n2+1)/2 // 2

    start, end := 0, n1

    for ; start <= end; {
        mid1 := (start+end)/2
        mid2 := left - mid1

        // fmt.Println("===============")
        // fmt.Println(mid1, ",", mid2)
        // fmt.Println(start, ",", end)
        if mid1 > start && nums1[mid1-1] > nums2[mid2] {
            end = mid1 - 1
        } else if mid1 < end && nums1[mid1] < nums2[mid2-1] {
            start = mid1 + 1
        } else {
            maxLeft := 0
            if mid1 == 0 {
                maxLeft = nums2[mid2-1]
            } else if mid2 == 0 {
                maxLeft = nums1[mid1-1]
            } else {
                maxLeft = max(nums1[mid1-1], nums2[mid2-1])
            }

            if (n1 + n2) % 2 == 1 {
                return float64(maxLeft)
            }

            minRight := 0
            if mid1 == n1 {
                minRight = nums2[mid2]
            } else if mid2 == n2 {
                minRight = nums1[mid1]
            } else {
                minRight = min(nums1[mid1], nums2[mid2])
            }

            return float64(maxLeft + minRight)/2.0
        }
    }

    return 0
}


func max(a , b int) int {
    if a > b {
        return a
    }

    return b
}

func min(a , b int) int {
    if a > b {
        return b
    }

    return a
}