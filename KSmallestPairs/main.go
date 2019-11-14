// KSmallestPairs
package main

func main() {

}


func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
    r := make([][]int, 0)
    if len(nums1) == 0 || len(nums2) == 0 {
        return r
    }

    order := true
    if len(nums1) > len(nums2) {
        nums2, nums1 = nums1, nums2
        order = false
    }

    p := make([]int, min(len(nums1), len(nums2)))

    t  := exchange(order, []int{nums1[0], nums2[p[0]]})
    p[0]++
    r = append(r, t)
    for x := 1; x < k; x++ {
        // 对于p里每一个元素, 需要去找寻最小的值
        v := 1000000000
        iv:= -1
        t = []int{0, 0}
        for i := 0; i < len(p) ; i++ {

            if p[i] >= len(nums2) {
                continue
            }

            if tv := nums1[i] + nums2[p[i]]; tv < v {
                iv = i
                v = tv
                t[0] = nums1[i]
                t[1] = nums2[p[i]]
                t = exchange(order, t)
            }
        }

        if v < 1000000000 {
            p[iv]++
            r = append(r, t)
        }
    }

    return r
}


func min(a , b int) int {
    if a > b {
        return b
    }

    return a
}

func exchange(order bool, a []int) []int{
    if order == false {
        a[0], a[1] = a[1], a[0]
    }

    return a
}