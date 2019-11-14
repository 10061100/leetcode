package main

import (
    "fmt"
    "sort"
)

func main() {
    // fmt.Println(threeSumMulti([]int{1,1,2,2,3,3,4,4,5,5}, 8) == 20 )
    // fmt.Println(threeSumMulti([]int{1,1,2,2,2,2}, 5) == 12)
    // fmt.Println(threeSumMulti([]int{1,2,3, 4,5,6,7,8,9}, 15) == 8)
    fmt.Println(threeSumMulti([]int{0,0,0,0, 0}, 0))
}

func threeSumMulti(A []int, target int) int {

    sort.Sort(sort.IntSlice(A))
    count := [101]int{}
    nums  := make([]int, 0)
    mod   := 100000007

    c := A[0]
    t := 1
    for i := 1; i < len(A); i++ {
        v := A[i]
        if v == c {
            t++
            continue
        }

        count[c] = t // 某个数字出现了n次
        nums = append(nums, c)

        t = 1
        c = v
    }

    count[c] = t
    nums = append(nums, c)

    res := 0
    // 统计单个可以进行3次计算
    for i := 0; i < len(nums); i++ {
        if count[nums[i]] < 3 {
            continue
        }

        if 3 * nums[i] != target {
            continue
        }

        t = count[nums[i]]

        var t1 = int64((t * (t-1)) % mod)
        t1 = (t1  * (int64(t)-int64(2))) % int64(mod)

        res += int(( t1 / 6) % int64(mod))
        res = res % mod
    }

    // 两个相同一个不同的 & 3个都不同的, 采用指针法
    for i := 0 ; i < len(nums) ; i++ {
        nt := target - nums[i]
        n  := count[nums[i]]
        // fmt.Println("======")
        // fmt.Println(nt)
        // 在 [i+1: ]之间找和等于nt的

        // 当前数一个, 右边2个
        t = 0
        for left, right := i+1, len(nums)-1; left <= right; {

            // 左边两个
            if left == right {
                if count[nums[left]] >= 2 && nums[left] * 2 == nt {
                    t += (count[nums[left]] * (count[nums[left]] - 1) / 2 * n) % mod
                    t = t % mod
                }
                break
            }

            if nums[left] * 2 > nt {
                break
            }

            if nums[right] * 2 < nt {
                break
            }

            if nums[left] + nums[right] < nt {
                left++
                continue
            }

            if nums[left] + nums[right] > nt {
                right--
                continue
            }

            // 左右各一个
            if nums[left] + nums[right] == nt {
                t += (count[nums[left]] * (count[nums[right]]) * n )%mod
                t = t % mod
            }

            left++
        }

        // fmt.Println(nt, ",", t)
        if count[nums[i]] >= 2 {
            nt = target - nums[i] * 2
            n  = count[nums[i]] * (count[nums[i]] - 1) /2
            for k := i+1; k < len(nums)  ; k++ {
                if nums[k] == nt {
                    t += (n * count[nums[k]]) % mod
                    t = t % mod
                }
            }
        }

        // fmt.Println(t)
        res  = ( res + t ) % mod
    }

    return res
}
