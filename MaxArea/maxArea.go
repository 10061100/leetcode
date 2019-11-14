// MaxArea
package main

func main() {
    print(maxArea([]int{1,2,1}))
}

func maxArea(height []int) int {
    length := len(height)

    if length < 2 {
        return 0
    }

    l := 0
    r := length - 1
    max := (r-l)*min(height[l], height[r])

    for ; l < r && l < length && r >= 0; {
        if height[l] >= height[r] {
            j := r - 1
            // 找到左边第一个比自己大的
            for ; j > l && height[j] <= height[r]; j-- {}
            if j <= l {
                break
            }

            if t := (j - l) * min(height[j], height[l]); t > max {
                max = t
            }

            r = j

        } else {
            j := l + 1
            for ; j < r && height[j] <= height[l]; j++{}
            if j >= r {
                break
            }

            if t := (r - j) * min(height[j], height[r]); t > max {
                max = t
            }
            l = j
        }
    }

    return max
}


func min(x, y int) int {
    if x > y {
        return y
    }
    return x
}
