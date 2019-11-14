// Merge Sorted Array
// MergeSortedArray
package main

import "fmt"

func main() {
    a := []int{1,2,3,0,0}
    b := []int{1,2}
    merge(a, 3, b, 2)
    fmt.Println(a)
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
    l1 := 0
    l2 := 0
    for ; ; {
        // println("l1:", l1, "l2:", l2)
        if l1 >= m + l2 {
            copy(nums1[l1:l1+n-l2], nums2[l2:n])
            l2 = n
        }

        if l2 >= n {

            break
        }

        step := 0
        for j := l2; j < n; j++ {
            if nums2[j] <= nums1[l1] {
                step++
            }
        }
        // l2 ~ l2 + step 都是 <= l1 的
        if step > 0 {
            arrayMoveBack(nums1, l2 + m, l1, step)
            copy(nums1[l1:l1+step], nums2[l2:l2+step])
            l1 += step
            l2 += step
        }

        l1++
    }
}


func arrayMoveBack(array []int, length, start, step int) {
    for  j := length - 1; j >= start; j-- {
        array[j+step] = array[j]
    }
}

