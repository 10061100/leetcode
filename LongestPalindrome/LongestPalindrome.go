
// LongestPalindrome
package main

func main() {
    println(longestPalindrome("bb"))
}

func longestPalindrome(s string) string {

    bs := []byte(s)
    result := ""
    if len(bs) == 0 {
        return result
    }

    max := 1
    result = string(bs[0:1])

    for i := 1; i < len(bs) ; i++ {
        c := 0
        if i < len(bs) - 1 && bs[i-1] == bs[i+1] {
            // 首先找那种两边对称的
            l, r := i-1, i+1
            c++
            for ; l >= 0 && r < len(bs) && bs[l] == bs[r]; {
                l--
                r++
                c += 2
            }
            if c > max {
                println(c)
                max = c
                result = string(bs[l+1:r])
            }
        }

        if bs[i-1] == bs[i] {
            // 再找那种无中心对称的
            l, r := i-1, i
            for ; l >= 0 && r < len(bs) && bs[l] == bs[r]; {
                l--
                r++
                c += 2
            }
            if c > max {
                max = c
                result = string(bs[l+1:r])
            }
        }

    }

    return result
}
