// CountBinarySubstrings

package main

func main() {
    // println(countBinarySubstrings("00110011") == 6)
    // println(countBinarySubstrings("10101") == 4)
    println(countBinarySubstrings("00110") == 3)
}

func countBinarySubstrings(s string) int {
    bs := []byte(s)
    cur := 1
    ans := 0
    pre := 0
    for i := 1; i < len(bs); i++ {
        println(cur)
        if bs[i] == bs[i-1] {
            cur++
        } else {
            if pre > 0 {
                ans += getBigger(pre, cur)
            }

            pre = cur
            cur = 1
        }
    }

    return ans + getBigger(pre, cur)
}


func getBigger(pre, cur int) int {
    if pre > cur {
        return cur
    }

    return pre
}